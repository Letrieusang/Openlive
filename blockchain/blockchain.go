package blockchain

import (
	"api-openlive/blockchain/factory"
	common2 "api-openlive/common"
	"api-openlive/config"
	marketstorage "api-openlive/module/market_place/storage"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	client      *ethclient.Client
	config      config.Blockchain
	marketStore marketstorage.MarketPlaceStore
}

type safeMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewClientHttp() *Client {
	config := config.GetBlockchainConfig()
	client, err := ethclient.Dial(config.GetHttpUrl())
	if err != nil {
		fmt.Println(err)
	}
	return &Client{
		client: client,
		config: config,
	}
}

func (c *Client) GetListItemOfOwner(Nfts []string, owner common.Address, itemContract common.Address) []*big.Int {
	var nftOfOwner []*big.Int
	fmt.Println(itemContract.Hex())
	itemFactory, err := factory.NewFactory(itemContract, c.client)
	if err != nil {
		log.Println("Create Item Factory error ", err.Error())
		return nftOfOwner
	}
	var wg sync.WaitGroup
	nftC := make(chan *big.Int, len(Nfts))
	for _, t := range Nfts {
		wg.Add(1)
		go func(nft string) {
			defer wg.Done()
			tBigInt := new(big.Int)
			tmp, ok := tBigInt.SetString(nft, 10)
			if !ok {
				log.Println("Invalid Nft: ", nft, "of owner: ", owner.Hex())
				return
			}
			result, err := itemFactory.OwnerOf(&bind.CallOpts{}, tmp)
			if err != nil {
				log.Println("Get Nft Error: ", nft, "of owner: ", owner.Hex(), "Error: ", err)
				return
			}
			if result.Hex() == owner.Hex() {
				nftC <- tmp
			}
		}(t)
	}
	wg.Wait()
	for len(nftC) > 0 {
		nftOfOwner = append(nftOfOwner, <-nftC)
	}
	return nftOfOwner
}

func (c *Client) GetNftURI(Nfts []*big.Int, itemContract common.Address) map[string]string {
	itemFactory, err := factory.NewFactory(itemContract, c.client)
	if err != nil {
		log.Println("Create Item Factory error %v", err.Error())
		return nil
	}
	var wg sync.WaitGroup
	mUri := safeMap{
		data: make(map[string]string),
	}
	for _, t := range Nfts {
		wg.Add(1)
		go func(nft *big.Int) {
			defer wg.Done()
			result, err := itemFactory.TokenURI(nil, nft)
			if err != nil {
				log.Println("Get Nft URI: ", t, "Error: ", err)
				return
			}
			fmt.Println("result: ", result)
			mUri.mu.Lock()
			mUri.data[nft.String()] = result
			mUri.mu.Unlock()
		}(t)
	}
	wg.Wait()
	return mUri.data
}

func (c *Client) GetNftURIOfUser(nfts []string, owner common.Address, itemContract common.Address) map[string]string {
	nftOfUser := c.GetListItemOfOwner(nfts, owner, itemContract)
	return c.GetNftURI(nftOfUser, itemContract)
}

func (c *Client) TransferToken(walletAddress string, amount float64) (string, error) {

	ctx := context.Background()
	nonce, err := c.client.PendingNonceAt(ctx, c.config.GetOurAddress())
	if err != nil {
		return "", err
	}
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	value := c.ConvertFloatToBigInt(amount)
	toAddress := common.HexToAddress(walletAddress)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	methodId := getMethodByType()
	if len(methodId) == 0 {
		return "", errors.New("Get Method id error")
	}
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)
	gasLimit, err := c.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     c.config.GetOurAddress(),
		To:       &toAddress,
		Data:     data,
		GasPrice: gasPrice,
	})
	fmt.Println("Gas limit: ", gasLimit)
	if err != nil {
		fmt.Println("???", err)
		return "", err
	}
	v := big.NewInt(0)
	tx := types.NewTransaction(nonce, c.config.GetOurAddress(), v, gasLimit*5, gasPrice, data)
	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.config.GetPrivate())
	if err != nil {
		return "", err
	}
	tx.RawSignatureValues()
	err = c.client.SendTransaction(context.Background(), signedTx)
	txHash := signedTx.Hash().Hex()
	if err != nil {
		return txHash, err
	}
	status := 0
	for i := 0; i <= 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		status = c.getTransaction(common.HexToHash(txHash))
		if status != common2.WITHDRAW_STATUS_PROCESSING {
			return txHash, nil
		}
	}
	return txHash, nil
}

func (c *Client) getTransaction(hash common.Hash) int {
	ctx := context.Background()
	rp, err := c.client.TransactionReceipt(ctx, hash)
	if err != nil || rp == nil {
		return common2.WITHDRAW_STATUS_PROCESSING
	}
	log.Printf("tx: %s, status: %v", hash.Hex(), rp.Status)
	if rp.Status == 1 {
		return common2.WITHDRAW_STATUS_APPROVED
	}
	if rp.Status == 0 {
		return common2.WITHDRAW_STATUS_FAILED
	}
	return common2.WITHDRAW_STATUS_PROCESSING
}

func getMethodByType() []byte {
	var name string = "transferFrom(uint256,address)"
	return getMethodId(name)
}

func getMethodId(name string) []byte {
	signature := []byte(name)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(signature)
	methodID := hash.Sum(nil)[:4]
	return methodID
}

func (c Client) ConvertFloatToBigInt(f float64) *big.Int {
	amount := strconv.FormatFloat(f, 'f', 17, 64)
	fmt.Println("Amount: ........", amount)
	nums := []rune(amount)
	result := []rune{}
	count := 0
	mark := false
	hasNumber := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == '.' {
			mark = true
			continue
		}
		if mark {
			count++
		}
		if (nums[i] == '0' && !hasNumber) || count > 18 {
			continue
		}
		result = append(result, nums[i])
		hasNumber = true
	}
	for i := 0; i < 18-count; i++ {
		result = append(result, '0')
	}
	n := new(big.Int)
	n.SetString(string(result), 10)
	return n
}
