package blockchain

import (
	"api-openlive/blockchain/transfer"
	common2 "api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	"context"
	"math/big"
)

func (s Synchronization) transferNft(log *transfer.TransferTransferNFT, ctx context.Context) {
	s.logger.Infof("Transfer Nft %v From %v To %v", log.TokenId.String(), log.Sender.Hex(), log.Receiver.Hex())
	users, err := s.userSvc.GetUsersByLogTransfer(log.Receiver, log.Sender)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqUpdate := marketplacemodel.UpdateItemOwnerRequest{}
	for _, u := range users {
		if u.WalletAddress == log.Sender.Hex() {
			reqUpdate.SenderId = u.Id
			continue
		}
		if u.WalletAddress == log.Receiver.Hex() {
			reqUpdate.ReceiverId = u.Id
			continue
		}
	}
	if reqUpdate.SenderId == 0 || reqUpdate.ReceiverId == 0 {
		s.logger.Errorf("Sender or receiver is invalid")
		return
	}
	reqUpdate.Nft = log.TokenId.String()
	reqUpdate.ContractAddress = log.ContractAddress.Hex()
	if err := s.marketSvc.TransferNft(reqUpdate); err != nil {
		s.logger.Errorf("Update owner: %v", err.Error())
		return
	}

	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Infof("Header Block Error: ", err.Error())
	}
	finalizeBlock.SetString(header.Number.String(), 10)
	s.logger.Infof("Header Block Number: ", header.Number)
	newBlock := marketplacemodel.BlockChainBlock{BlockNumber: finalizeBlock.Uint64(), Status: common2.BLOCK_STATUS_WS_SUCCESS, IsBlockProcess: common2.IS_NOT_BLOCK_PROCESS}
	errCreateBlock := s.marketSvc.Store.AddWsBlock(newBlock)
	if errCreateBlock != nil {
		s.logger.Infof("Create Block Error: ", errCreateBlock.Error())
	}
	s.logger.Infof("Update Owner successfully")
}
