package config

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"log"
	"os"
)

var config *viper.Viper

type Blockchain struct {
	httpUrl                       string
	wsUrl                         string
	privateKey                    *ecdsa.PrivateKey
	publicKey                     *ecdsa.PublicKey
	publicAddress                 common.Address
	ourAddress                    common.Address
	opvPoolAddress                common.Address
	opvMarketplaceAddressFirst    common.Address
	opvMarketplaceAddressSecond   common.Address
	opvMarketplaceAddressThird    common.Address
	opvFactoryAddress             common.Address
	opvNFTTransferAddress         common.Address
	opvNFTCreateEventAddress      common.Address
	opvNFTCreateEventSecAddress   common.Address
	opvCreateNftAddress           common.Address
	opvAuctionAddress             common.Address
	opvMarketplaceBuyTopic        common.Hash
	opvMarketplaceSellTopic       common.Hash
	opvMarketplaceCancelSellTopic common.Hash
	opvNFTTransferTopic           common.Hash
	opvNFTCreateEventTopic        common.Hash
	opvNFTCreateEventSecTopic     common.Hash
	opvNFTCancelEventTopic        common.Hash
	opvNFTCancelEventSecTopic     common.Hash
	opvNFTMintTopic               common.Hash
	opvAuctionTopic               common.Hash
	opvAuctionBidTopic            common.Hash
	opvAuctionClaimNftTopic       common.Hash
	opvAuctionCancelTopic         common.Hash
	opvTransferTopic              common.Hash
	opvMultiTransferTopic         common.Hash
	opvNewPackTopic               common.Hash
	opvNewPackAuctionTopic        common.Hash
	opvPackCancelSaleTopic        common.Hash
	opvPackCancelAuctionTopic     common.Hash
}

type Server struct {
	port       int
	sqlPath    string
	rdsAddr    string
	rdsPwd     string
	privateKey string
	sentry     string
	apiDomain  string
}

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	fmt.Println("env", env)
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	os.Setenv("env", env)
}

func GetBlockchainConfig() Blockchain {
	var b Blockchain
	privateKey, err := crypto.HexToECDSA(config.GetString("blockchain.private_key"))
	if err != nil {
		log.Fatalln("Can not get private key ", err)
	}
	b.privateKey = privateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("Can not get public key ", err)
	}
	b.publicKey = publicKeyECDSA
	b.ourAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	b.httpUrl = config.GetString("blockchain.http_url")
	b.wsUrl = config.GetString("blockchain.ws_url")
	b.publicAddress = common.HexToAddress(config.GetString("blockchain.public_address"))
	b.opvPoolAddress = common.HexToAddress(config.GetString("blockchain.opv_pool_address"))
	b.opvMarketplaceAddressFirst = common.HexToAddress(config.GetString("blockchain.opv_marketplace_address_first"))
	b.opvFactoryAddress = common.HexToAddress(config.GetString("blockchain.opv_factory_address"))
	b.opvMarketplaceAddressSecond = common.HexToAddress(config.GetString("blockchain.opv_marketplace_address_second"))
	b.opvMarketplaceAddressThird = common.HexToAddress(config.GetString("blockchain.opv_marketplace_address_third"))
	b.opvNFTTransferAddress = common.HexToAddress(config.GetString("blockchain.opv_nft_transfer_address"))
	b.opvNFTCreateEventAddress = common.HexToAddress(config.GetString("blockchain.opv_nft_create_event_address"))
	b.opvNFTCreateEventSecAddress = common.HexToAddress(config.GetString("blockchain.opv_nft_create_event_sec_address"))
	b.opvCreateNftAddress = common.HexToAddress(config.GetString("blockchain.opv_create_nft_address"))
	b.opvAuctionAddress = common.HexToAddress(config.GetString("blockchain.opv_auction_address"))
	b.opvMarketplaceBuyTopic = common.HexToHash(config.GetString("blockchain.opv_marketplace_buy_topic"))
	b.opvMarketplaceSellTopic = common.HexToHash(config.GetString("blockchain.opv_marketplace_sell_topic"))
	b.opvMarketplaceCancelSellTopic = common.HexToHash(config.GetString("blockchain.opv_marketplace_cancel_sell_topic"))
	b.opvNFTTransferTopic = common.HexToHash(config.GetString("blockchain.opv_nft_transfer_topic"))
	b.opvNFTCreateEventTopic = common.HexToHash(config.GetString("blockchain.opv_nft_create_event_topic"))
	b.opvNFTCreateEventSecTopic = common.HexToHash(config.GetString("blockchain.opv_nft_create_event_sec_topic"))
	b.opvNFTCancelEventTopic = common.HexToHash(config.GetString("blockchain.opv_nft_cancel_event_topic"))
	b.opvNFTCancelEventSecTopic = common.HexToHash(config.GetString("blockchain.opv_nft_cancel_event_sec_topic"))
	b.opvNFTMintTopic = common.HexToHash(config.GetString("blockchain.opv_nft_mint_topic"))
	b.opvAuctionTopic = common.HexToHash(config.GetString("blockchain.opv_auction_topic"))
	b.opvAuctionBidTopic = common.HexToHash(config.GetString("blockchain.opv_auction_bid_topic"))
	b.opvAuctionClaimNftTopic = common.HexToHash(config.GetString("blockchain.opv_auction_claim_nft_topic"))
	b.opvAuctionCancelTopic = common.HexToHash(config.GetString("blockchain.opv_auction_cancel_topic"))
	b.opvTransferTopic = common.HexToHash(config.GetString("blockchain.opv_transfer_topic"))
	b.opvMultiTransferTopic = common.HexToHash(config.GetString("blockchain.opv_multi_transfer_topic"))
	b.opvNewPackTopic = common.HexToHash(config.GetString("blockchain.opv_new_pack_topic"))
	b.opvNewPackAuctionTopic = common.HexToHash(config.GetString("blockchain.opv_new_pack_auction_topic"))
	b.opvPackCancelAuctionTopic = common.HexToHash(config.GetString("blockchain.opv_pack_cancel_auction_topic"))
	b.opvPackCancelSaleTopic = common.HexToHash(config.GetString("blockchain.opv_pack_cancel_sale_topic"))
	return b
}

func GetServer() Server {
	var s Server
	s.port = config.GetInt("server.port")
	s.apiDomain = config.GetString("server.api_domain")
	s.sentry = config.GetString("server.sentry")
	s.sqlPath = config.GetString("server.sql_path")
	s.privateKey = config.GetString("server.private_key")
	s.rdsAddr = config.GetString("server.rds_addr")
	s.rdsPwd = config.GetString("server.rds_pwd")
	return s
}

func (b Blockchain) GetHttpUrl() string {
	return b.httpUrl
}

func (b Blockchain) GetOpvMarketplaceAddressFirst() common.Address {
	return b.opvMarketplaceAddressFirst
}

func (b Blockchain) GetOpvMarketplaceAddressSecond() common.Address {
	return b.opvMarketplaceAddressSecond
}

func (b Blockchain) GetOpvMarketplaceAddressThird() common.Address {
	return b.opvMarketplaceAddressThird
}

func (b Blockchain) GetOpvFactoryAddress() common.Address {
	return b.opvFactoryAddress
}

func (b Blockchain) GetOpvCreateNftAddress() common.Address {
	return b.opvCreateNftAddress
}

func (b Blockchain) GetWsUrl() string {
	return b.wsUrl
}

func (b Blockchain) GetOpvNftTransferAddress() common.Address {
	return b.opvNFTTransferAddress
}

func (b Blockchain) GetOpvNftCreateEventAddress() common.Address {
	return b.opvNFTCreateEventAddress
}

func (b Blockchain) GetOpvNftCreateEventSecAddress() common.Address {
	return b.opvNFTCreateEventSecAddress
}

func (b Blockchain) GetOpvMarketplaceBuyTopic() common.Hash {
	return b.opvMarketplaceBuyTopic
}
func (b Blockchain) GetOpvMarketplaceSellTopic() common.Hash {
	return b.opvMarketplaceSellTopic
}
func (b Blockchain) GetOpvMarketplaceCancelSellTopic() common.Hash {
	return b.opvMarketplaceCancelSellTopic
}
func (b Blockchain) GetOpvNFTTransferTopic() common.Hash {
	return b.opvNFTTransferTopic
}

func (b Blockchain) GetOpvNftCreateEventTopic() common.Hash {
	return b.opvNFTCreateEventTopic
}

func (b Blockchain) GetOpvNFTCancelEventTopic() common.Hash {
	return b.opvNFTCancelEventTopic
}

func (b Blockchain) GetOpvNftCreateEventSecTopic() common.Hash {
	return b.opvNFTCreateEventSecTopic
}

func (b Blockchain) GetOpvNFTCancelEventSecTopic() common.Hash {
	return b.opvNFTCancelEventSecTopic
}

func (b Blockchain) GetOpvNftMintTopic() common.Hash {
	return b.opvNFTMintTopic
}

func (b Blockchain) GetPrivate() *ecdsa.PrivateKey {
	return b.privateKey
}

func (b Blockchain) GetOurAddress() common.Address {
	return b.ourAddress
}

func (b Blockchain) GetPublicAddress() common.Address {
	return b.publicAddress
}

func (b Blockchain) GetOpvPoolAddress() *common.Address {
	return &b.opvPoolAddress
}

func (b Blockchain) GetOpvPoolAddressSec() common.Address {
	return b.opvPoolAddress
}

func (b Blockchain) GetOpvAuctionAddress() common.Address {
	return b.opvAuctionAddress
}

func (b Blockchain) GetOpvAuctionTopic() common.Hash {
	return b.opvAuctionTopic
}

func (b Blockchain) GetOpvAuctionBidTopic() common.Hash {
	return b.opvAuctionBidTopic
}

func (b Blockchain) GetOpvAuctionClaimNftTopic() common.Hash {
	return b.opvAuctionClaimNftTopic
}

func (b Blockchain) GetOpvAuctionCancelTopic() common.Hash {
	return b.opvAuctionCancelTopic
}

func (b Blockchain) GetOpvNewPackTopic() common.Hash {
	return b.opvNewPackTopic
}

func (b Blockchain) GetOpvNewPackAuctionTopic() common.Hash {
	return b.opvNewPackAuctionTopic
}

func (b Blockchain) GetOpvPackCancelSaleTopic() common.Hash {
	return b.opvPackCancelSaleTopic
}

func (b Blockchain) GetOpvPackCancelAuctionTopic() common.Hash {
	return b.opvPackCancelAuctionTopic
}

func (b Blockchain) GetOpvTransferTopic() common.Hash {
	return b.opvTransferTopic
}

func (b Blockchain) GetOpvMultiTransferTopic() common.Hash {
	return b.opvMultiTransferTopic
}

func (s Server) GetApiDomain() string {
	return s.apiDomain
}

func GetConfig() *viper.Viper {
	return config
}
