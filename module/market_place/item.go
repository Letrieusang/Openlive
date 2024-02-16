package market_place

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	"encoding/json"
	"fmt"
	common_eth "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
)

type detailItem struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d detailItem) convertToItem(userId int64, nft string, contractAddress string) marketplacemodel.Item {
	return marketplacemodel.Item{
		UserId:          userId,
		ContractAddress: contractAddress,
		Title:           d.Name,
		Nft:             nft,
		Status:          common.ITEM_NEW,
		Image:           d.Image,
		Description:     d.Description,
	}
}

// @Sumary Sync Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags Item
// @Accept json
// @Produce json
// @Param user body marketplacemodel.ItemWithContract true "sync items request"
// @Success 200 {object} common.Response{data=marketplacemodel.CreateItemResult}
// @Failure 400,401,500 {object} common.Response
// @Router /item/sync [post]
func (u MarketPlaceHandler) SyncItems() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.ItemWithContract
		if res := ctx.BindJSON(&rq); res != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": res.Error(),
			})
			return
		}

		itemContract := common_eth.HexToAddress(rq.ContractAddress)

		if len(rq.NftItems) > 50 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "The length of array must smaller than 50",
			})
			return
		}

		userId := ctx.GetInt("id")
		rq.ContractAddress = itemContract.Hex()
		walletAddress, nfts, err := u.Client.GetNftsAndWalletAddress(rq, int64(userId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		mNftUri := u.Blc.GetNftURIOfUser(nfts, common_eth.HexToAddress(walletAddress), itemContract) // get URI of NFT
		items := u.getItemByURI(mNftUri, int64(userId), itemContract.Hex())
		var totalItem int64
		if len(items) == 0 {
			totalItem = 0
		} else {
			totalItem, err = u.Client.BatchInsertNft(items)
		}

		if err != nil {
			ctx.JSON(http.StatusOK, common.Response{
				HasError: 1,
				Message:  "Sync Item Fail",
				Data:     err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.Response{
			HasError: 0,
			Message:  "Sync Item result",
			Data: marketplacemodel.CreateItemResult{
				TotalItemCreated: totalItem,
				ContractAddress:  rq.ContractAddress,
			},
		})

	}
}

func (u MarketPlaceHandler) getItemByURI(mNftUri map[string]string, userId int64, contractAddress string) []marketplacemodel.Item {
	var items []marketplacemodel.Item
	itemC := make(chan marketplacemodel.Item, len(mNftUri))
	var wg sync.WaitGroup
	for k, v := range mNftUri {
		wg.Add(1)
		go func(nft string, uri string) {
			defer wg.Done()
			var tmp detailItem
			resp, err := http.Get(uri)
			if err != nil {
				itemC <- marketplacemodel.DefaultItem(userId, nft, contractAddress)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				itemC <- marketplacemodel.DefaultItem(userId, nft, contractAddress)
				return
			}
			json.Unmarshal(body, &tmp)
			if tmp.Name != "" {
				itemC <- tmp.convertToItem(userId, nft, contractAddress)
			} else {
				itemC <- marketplacemodel.DefaultItem(userId, nft, contractAddress)
			}
		}(k, v)
	}
	wg.Wait()
	for len(itemC) > 0 {
		items = append(items, <-itemC)
	}
	fmt.Println(items)
	return items
}
