package business

import (
	commonerr "api-openlive/common"
	"api-openlive/config"
	marketplacemodel "api-openlive/module/market_place/model"
	"fmt"
)

type detailItem struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d detailItem) convertToItem(userId int64, nft string) marketplacemodel.Item {
	return marketplacemodel.Item{
		UserId:          userId,
		ContractAddress: config.GetBlockchainConfig().GetOpvFactoryAddress().Hex(),
		Title:           d.Name,
		Nft:             nft,
		Status:          commonerr.ITEM_NEW,
		Image:           d.Image,
		Description:     d.Description,
	}
}

func (i MarketPlaceController) GetNftsAndWalletAddress(rq marketplacemodel.ItemWithContract, id int64) (string, []string, error) {
	var reqNfts []string
	var nfts []string
	mnfts := make(map[string]bool)
	for _, t := range rq.NftItems {
		if !mnfts[t.Id] {
			reqNfts = append(reqNfts, t.Id)
			mnfts[t.Id] = true
		}
	}
	mIds, err := i.Store.GetByUserIdAndNftIds(id, reqNfts, rq.ContractAddress)
	if err != nil {
		return "", nfts, err
	}

	for _, t := range rq.NftItems {
		if !mIds[t.Id] {
			nfts = append(nfts, t.Id)
		}
		mIds[t.Id] = true
	}
	fmt.Println(nfts)
	walletAddress, err := i.UserStore.GetWalletAddressById(id)
	if err != nil {
		return "", nfts, commonerr.ErrFindWalletAddress
	}
	return walletAddress, nfts, nil
}

func (i MarketPlaceController) BatchInsertNft(items []marketplacemodel.Item) (int64, error) {
	return i.Store.BatchInsertItem(items)
}
