// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace_multi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LibMarketplaceStorageMarketItem is an auto generated low-level Go binding around an user-defined struct.
type LibMarketplaceStorageMarketItem struct {
	IdOnMarket   *big.Int
	TokenId      *big.Int
	ContractHold common.Address
	NftContract  common.Address
	Seller       common.Address
	Owner        common.Address
	Price        *big.Int
	Sold         bool
	PackId       *big.Int
}

// LibMarketplaceStoragePackData is an auto generated low-level Go binding around an user-defined struct.
type LibMarketplaceStoragePackData struct {
	IsPackSold bool
	FromId     *big.Int
	EndId      *big.Int
}

// MarketplaceMultiMetaData contains all meta data concerning the MarketplaceMulti contract.
var MarketplaceMultiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CancelSellAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ClaimNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"}],\"name\":\"CreateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnBid\",\"type\":\"uint256\"}],\"name\":\"LoseBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"storageIdOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"NewAuctionPack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelSellAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchClaimNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"batchSaleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"cancelSellAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"claimNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getListingItemInOnePackAuction\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"saleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyMarketItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CancelSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"CreateMarketItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"NewPack\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceItem\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"}],\"name\":\"Sale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idsOnMarket\",\"type\":\"uint256[]\"},{\"internalType\":\"enumIMarketplaceFeature.BATCH_BUY_TYPE\",\"name\":\"buyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"batchBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceItem\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"}],\"name\":\"batchSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"cancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIdOnMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"}],\"internalType\":\"structLibMarketplaceStorage.MarketItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getIsPackSold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getListingItemInOnePack\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"}],\"name\":\"getPack\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPackSold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fromId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"}],\"internalType\":\"structLibMarketplaceStorage.PackData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"AddBlackListFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"AddFactoryBasic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"AddFactoryVip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"NewFeeAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeRef\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mainToken\",\"type\":\"address\"}],\"name\":\"NewMainTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refAddress\",\"type\":\"address\"}],\"name\":\"NewRefAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"RemoveBlackListFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RemoveFactoryBasic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RemoveFactoryVip\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkBlackListFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"checkFactoryBasic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"checkFactoryVip\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getBuyData\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeCreator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRef\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRefAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"removeBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"setBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeRef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MainToken\",\"type\":\"address\"}],\"name\":\"setMainTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_RefAddress\",\"type\":\"address\"}],\"name\":\"setRefAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceMultiABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMultiMetaData.ABI instead.
var MarketplaceMultiABI = MarketplaceMultiMetaData.ABI

// MarketplaceMulti is an auto generated Go binding around an Ethereum contract.
type MarketplaceMulti struct {
	MarketplaceMultiCaller     // Read-only binding to the contract
	MarketplaceMultiTransactor // Write-only binding to the contract
	MarketplaceMultiFilterer   // Log filterer for contract events
}

// MarketplaceMultiCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceMultiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceMultiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceMultiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceMultiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceMultiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceMultiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceMultiSession struct {
	Contract     *MarketplaceMulti // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceMultiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceMultiCallerSession struct {
	Contract *MarketplaceMultiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MarketplaceMultiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceMultiTransactorSession struct {
	Contract     *MarketplaceMultiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MarketplaceMultiRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceMultiRaw struct {
	Contract *MarketplaceMulti // Generic contract binding to access the raw methods on
}

// MarketplaceMultiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceMultiCallerRaw struct {
	Contract *MarketplaceMultiCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceMultiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceMultiTransactorRaw struct {
	Contract *MarketplaceMultiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplaceMulti creates a new instance of MarketplaceMulti, bound to a specific deployed contract.
func NewMarketplaceMulti(address common.Address, backend bind.ContractBackend) (*MarketplaceMulti, error) {
	contract, err := bindMarketplaceMulti(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMulti{MarketplaceMultiCaller: MarketplaceMultiCaller{contract: contract}, MarketplaceMultiTransactor: MarketplaceMultiTransactor{contract: contract}, MarketplaceMultiFilterer: MarketplaceMultiFilterer{contract: contract}}, nil
}

// NewMarketplaceMultiCaller creates a new read-only instance of MarketplaceMulti, bound to a specific deployed contract.
func NewMarketplaceMultiCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceMultiCaller, error) {
	contract, err := bindMarketplaceMulti(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiCaller{contract: contract}, nil
}

// NewMarketplaceMultiTransactor creates a new write-only instance of MarketplaceMulti, bound to a specific deployed contract.
func NewMarketplaceMultiTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceMultiTransactor, error) {
	contract, err := bindMarketplaceMulti(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiTransactor{contract: contract}, nil
}

// NewMarketplaceMultiFilterer creates a new log filterer instance of MarketplaceMulti, bound to a specific deployed contract.
func NewMarketplaceMultiFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceMultiFilterer, error) {
	contract, err := bindMarketplaceMulti(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiFilterer{contract: contract}, nil
}

// bindMarketplaceMulti binds a generic wrapper to an already deployed contract.
func bindMarketplaceMulti(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceMultiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketplaceMulti *MarketplaceMultiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketplaceMulti.Contract.MarketplaceMultiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketplaceMulti *MarketplaceMultiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.MarketplaceMultiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketplaceMulti *MarketplaceMultiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.MarketplaceMultiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketplaceMulti *MarketplaceMultiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketplaceMulti.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketplaceMulti *MarketplaceMultiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketplaceMulti *MarketplaceMultiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.contract.Transact(opts, method, params...)
}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCaller) CheckBlackListFee(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "checkBlackListFee", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiSession) CheckBlackListFee(user common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckBlackListFee(&_MarketplaceMulti.CallOpts, user)
}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) CheckBlackListFee(user common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckBlackListFee(&_MarketplaceMulti.CallOpts, user)
}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCaller) CheckFactoryBasic(opts *bind.CallOpts, proxy common.Address) (bool, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "checkFactoryBasic", proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiSession) CheckFactoryBasic(proxy common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckFactoryBasic(&_MarketplaceMulti.CallOpts, proxy)
}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) CheckFactoryBasic(proxy common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckFactoryBasic(&_MarketplaceMulti.CallOpts, proxy)
}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCaller) CheckFactoryVip(opts *bind.CallOpts, proxy common.Address) (bool, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "checkFactoryVip", proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiSession) CheckFactoryVip(proxy common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckFactoryVip(&_MarketplaceMulti.CallOpts, proxy)
}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) CheckFactoryVip(proxy common.Address) (bool, error) {
	return _MarketplaceMulti.Contract.CheckFactoryVip(&_MarketplaceMulti.CallOpts, proxy)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_MarketplaceMulti *MarketplaceMultiCaller) GetBuyData(opts *bind.CallOpts, price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getBuyData", price_, factory_, seller_, tokenId_)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_MarketplaceMulti *MarketplaceMultiSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _MarketplaceMulti.Contract.GetBuyData(&_MarketplaceMulti.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _MarketplaceMulti.Contract.GetBuyData(&_MarketplaceMulti.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiSession) GetFeeAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetFeeAddress(&_MarketplaceMulti.CallOpts)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetFeeAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetFeeAddress(&_MarketplaceMulti.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetFeeCreator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getFeeCreator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiSession) GetFeeCreator() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeCreator(&_MarketplaceMulti.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetFeeCreator() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeCreator(&_MarketplaceMulti.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetFeeMarket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getFeeMarket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiSession) GetFeeMarket() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeMarket(&_MarketplaceMulti.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetFeeMarket() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeMarket(&_MarketplaceMulti.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetFeeRef(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getFeeRef")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiSession) GetFeeRef() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeRef(&_MarketplaceMulti.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetFeeRef() (*big.Int, error) {
	return _MarketplaceMulti.Contract.GetFeeRef(&_MarketplaceMulti.CallOpts)
}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_MarketplaceMulti *MarketplaceMultiCaller) GetIdOnMarket(opts *bind.CallOpts, id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getIdOnMarket", id)

	if err != nil {
		return *new(LibMarketplaceStorageMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new(LibMarketplaceStorageMarketItem)).(*LibMarketplaceStorageMarketItem)

	return out0, err

}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_MarketplaceMulti *MarketplaceMultiSession) GetIdOnMarket(id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	return _MarketplaceMulti.Contract.GetIdOnMarket(&_MarketplaceMulti.CallOpts, id)
}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetIdOnMarket(id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	return _MarketplaceMulti.Contract.GetIdOnMarket(&_MarketplaceMulti.CallOpts, id)
}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetIsPackSold(opts *bind.CallOpts, _pack *big.Int) (bool, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getIsPackSold", _pack)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiSession) GetIsPackSold(_pack *big.Int) (bool, error) {
	return _MarketplaceMulti.Contract.GetIsPackSold(&_MarketplaceMulti.CallOpts, _pack)
}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetIsPackSold(_pack *big.Int) (bool, error) {
	return _MarketplaceMulti.Contract.GetIsPackSold(&_MarketplaceMulti.CallOpts, _pack)
}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiCaller) GetListingItemInOnePack(opts *bind.CallOpts, _pack *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getListingItemInOnePack", _pack)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiSession) GetListingItemInOnePack(_pack *big.Int) ([]*big.Int, error) {
	return _MarketplaceMulti.Contract.GetListingItemInOnePack(&_MarketplaceMulti.CallOpts, _pack)
}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetListingItemInOnePack(_pack *big.Int) ([]*big.Int, error) {
	return _MarketplaceMulti.Contract.GetListingItemInOnePack(&_MarketplaceMulti.CallOpts, _pack)
}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiCaller) GetListingItemInOnePackAuction(opts *bind.CallOpts, _pack *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getListingItemInOnePackAuction", _pack)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiSession) GetListingItemInOnePackAuction(_pack *big.Int) ([]*big.Int, error) {
	return _MarketplaceMulti.Contract.GetListingItemInOnePackAuction(&_MarketplaceMulti.CallOpts, _pack)
}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetListingItemInOnePackAuction(_pack *big.Int) ([]*big.Int, error) {
	return _MarketplaceMulti.Contract.GetListingItemInOnePackAuction(&_MarketplaceMulti.CallOpts, _pack)
}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetMainTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getMainTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiSession) GetMainTokenAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetMainTokenAddress(&_MarketplaceMulti.CallOpts)
}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetMainTokenAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetMainTokenAddress(&_MarketplaceMulti.CallOpts)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_MarketplaceMulti *MarketplaceMultiCaller) GetPack(opts *bind.CallOpts, pack *big.Int) (LibMarketplaceStoragePackData, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getPack", pack)

	if err != nil {
		return *new(LibMarketplaceStoragePackData), err
	}

	out0 := *abi.ConvertType(out[0], new(LibMarketplaceStoragePackData)).(*LibMarketplaceStoragePackData)

	return out0, err

}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_MarketplaceMulti *MarketplaceMultiSession) GetPack(pack *big.Int) (LibMarketplaceStoragePackData, error) {
	return _MarketplaceMulti.Contract.GetPack(&_MarketplaceMulti.CallOpts, pack)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetPack(pack *big.Int) (LibMarketplaceStoragePackData, error) {
	return _MarketplaceMulti.Contract.GetPack(&_MarketplaceMulti.CallOpts, pack)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCaller) GetRefAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketplaceMulti.contract.Call(opts, &out, "getRefAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiSession) GetRefAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetRefAddress(&_MarketplaceMulti.CallOpts)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_MarketplaceMulti *MarketplaceMultiCallerSession) GetRefAddress() (common.Address, error) {
	return _MarketplaceMulti.Contract.GetRefAddress(&_MarketplaceMulti.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) Bid(opts *bind.TransactOpts, _idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "Bid", _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Bid(&_MarketplaceMulti.TransactOpts, _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Bid(&_MarketplaceMulti.TransactOpts, _idOnMarket, _price)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) Buy(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "Buy", _idOnMarket)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) Buy(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Buy(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Buy(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Buy(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) Sale(opts *bind.TransactOpts, _tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "Sale", _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) Sale(_tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Sale(&_MarketplaceMulti.TransactOpts, _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Sale(_tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Sale(&_MarketplaceMulti.TransactOpts, _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) AddFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "addFactoryBasic", proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.AddFactoryBasic(&_MarketplaceMulti.TransactOpts, proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.AddFactoryBasic(&_MarketplaceMulti.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) AddFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "addFactoryVip", proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.AddFactoryVip(&_MarketplaceMulti.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.AddFactoryVip(&_MarketplaceMulti.TransactOpts, proxy)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchBuy(opts *bind.TransactOpts, packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchBuy", packId, _idsOnMarket, buyType, _price)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchBuy(packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchBuy(&_MarketplaceMulti.TransactOpts, packId, _idsOnMarket, buyType, _price)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchBuy(packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchBuy(&_MarketplaceMulti.TransactOpts, packId, _idsOnMarket, buyType, _price)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchCancelSell(opts *bind.TransactOpts, packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchCancelSell", packId, _idOnMarket)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchCancelSell(packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchCancelSell(&_MarketplaceMulti.TransactOpts, packId, _idOnMarket)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchCancelSell(packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchCancelSell(&_MarketplaceMulti.TransactOpts, packId, _idOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchCancelSellAuction(opts *bind.TransactOpts, packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchCancelSellAuction", packId, idsOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchCancelSellAuction(packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchCancelSellAuction(&_MarketplaceMulti.TransactOpts, packId, idsOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchCancelSellAuction(packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchCancelSellAuction(&_MarketplaceMulti.TransactOpts, packId, idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchClaimNft(opts *bind.TransactOpts, pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchClaimNft", pack, _idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchClaimNft(pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchClaimNft(&_MarketplaceMulti.TransactOpts, pack, _idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchClaimNft(pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchClaimNft(&_MarketplaceMulti.TransactOpts, pack, _idsOnMarket)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchSale(opts *bind.TransactOpts, _tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchSale", _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchSale(_tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchSale(&_MarketplaceMulti.TransactOpts, _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchSale(_tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchSale(&_MarketplaceMulti.TransactOpts, _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) BatchSaleAuction(opts *bind.TransactOpts, _tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "batchSaleAuction", _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) BatchSaleAuction(_tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchSaleAuction(&_MarketplaceMulti.TransactOpts, _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) BatchSaleAuction(_tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.BatchSaleAuction(&_MarketplaceMulti.TransactOpts, _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) CancelSell(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "cancelSell", _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.CancelSell(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.CancelSell(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) CancelSellAuction(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "cancelSellAuction", _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) CancelSellAuction(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.CancelSellAuction(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) CancelSellAuction(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.CancelSellAuction(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) ClaimNft(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "claimNft", _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.ClaimNft(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.ClaimNft(&_MarketplaceMulti.TransactOpts, _idOnMarket)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiSession) Migrate() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate(&_MarketplaceMulti.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Migrate() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate(&_MarketplaceMulti.TransactOpts)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactor) Migrate0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "migrate0")
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiSession) Migrate0() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate0(&_MarketplaceMulti.TransactOpts)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Migrate0() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate0(&_MarketplaceMulti.TransactOpts)
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactor) Migrate1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "migrate1")
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiSession) Migrate1() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate1(&_MarketplaceMulti.TransactOpts)
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) Migrate1() (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.Migrate1(&_MarketplaceMulti.TransactOpts)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) RemoveBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "removeBlackListFee", user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveBlackListFee(&_MarketplaceMulti.TransactOpts, user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveBlackListFee(&_MarketplaceMulti.TransactOpts, user)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) RemoveFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "removeFactoryBasic", proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveFactoryBasic(&_MarketplaceMulti.TransactOpts, proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveFactoryBasic(&_MarketplaceMulti.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) RemoveFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "removeFactoryVip", proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveFactoryVip(&_MarketplaceMulti.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.RemoveFactoryVip(&_MarketplaceMulti.TransactOpts, proxy)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SaleAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "saleAuction", _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SaleAuction(&_MarketplaceMulti.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SaleAuction(&_MarketplaceMulti.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setBlackListFee", user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetBlackListFee(&_MarketplaceMulti.TransactOpts, user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetBlackListFee(&_MarketplaceMulti.TransactOpts, user)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeAddress(&_MarketplaceMulti.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeAddress(&_MarketplaceMulti.TransactOpts, _feeAddress)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetFeeCreator(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setFeeCreator", percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeCreator(&_MarketplaceMulti.TransactOpts, percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeCreator(&_MarketplaceMulti.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetFeeMarket(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setFeeMarket", percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeMarket(&_MarketplaceMulti.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeMarket(&_MarketplaceMulti.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetFeeRef(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setFeeRef", percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeRef(&_MarketplaceMulti.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetFeeRef(&_MarketplaceMulti.TransactOpts, percent)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetMainTokenAddress(opts *bind.TransactOpts, _MainToken common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setMainTokenAddress", _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetMainTokenAddress(&_MarketplaceMulti.TransactOpts, _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetMainTokenAddress(&_MarketplaceMulti.TransactOpts, _MainToken)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactor) SetRefAddress(opts *bind.TransactOpts, _RefAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.contract.Transact(opts, "setRefAddress", _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetRefAddress(&_MarketplaceMulti.TransactOpts, _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_MarketplaceMulti *MarketplaceMultiTransactorSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _MarketplaceMulti.Contract.SetRefAddress(&_MarketplaceMulti.TransactOpts, _RefAddress)
}

// MarketplaceMultiAddBlackListFeeIterator is returned from FilterAddBlackListFee and is used to iterate over the raw logs and unpacked data for AddBlackListFee events raised by the MarketplaceMulti contract.
type MarketplaceMultiAddBlackListFeeIterator struct {
	Event *MarketplaceMultiAddBlackListFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiAddBlackListFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiAddBlackListFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiAddBlackListFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiAddBlackListFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiAddBlackListFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiAddBlackListFee represents a AddBlackListFee event raised by the MarketplaceMulti contract.
type MarketplaceMultiAddBlackListFee struct {
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddBlackListFee is a free log retrieval operation binding the contract event 0x6c23ff28a8bb18a73e79a449be6a4435f5123b723715d33899f0b39e2b2beb07.
//
// Solidity: event AddBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterAddBlackListFee(opts *bind.FilterOpts) (*MarketplaceMultiAddBlackListFeeIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "AddBlackListFee")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiAddBlackListFeeIterator{contract: _MarketplaceMulti.contract, event: "AddBlackListFee", logs: logs, sub: sub}, nil
}

// WatchAddBlackListFee is a free log subscription operation binding the contract event 0x6c23ff28a8bb18a73e79a449be6a4435f5123b723715d33899f0b39e2b2beb07.
//
// Solidity: event AddBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchAddBlackListFee(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiAddBlackListFee) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "AddBlackListFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiAddBlackListFee)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "AddBlackListFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddBlackListFee is a log parse operation binding the contract event 0x6c23ff28a8bb18a73e79a449be6a4435f5123b723715d33899f0b39e2b2beb07.
//
// Solidity: event AddBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseAddBlackListFee(log types.Log) (*MarketplaceMultiAddBlackListFee, error) {
	event := new(MarketplaceMultiAddBlackListFee)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "AddBlackListFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiAddFactoryBasicIterator is returned from FilterAddFactoryBasic and is used to iterate over the raw logs and unpacked data for AddFactoryBasic events raised by the MarketplaceMulti contract.
type MarketplaceMultiAddFactoryBasicIterator struct {
	Event *MarketplaceMultiAddFactoryBasic // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiAddFactoryBasicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiAddFactoryBasic)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiAddFactoryBasic)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiAddFactoryBasicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiAddFactoryBasicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiAddFactoryBasic represents a AddFactoryBasic event raised by the MarketplaceMulti contract.
type MarketplaceMultiAddFactoryBasic struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddFactoryBasic is a free log retrieval operation binding the contract event 0xdc6ec69d793d0fab8b57c062c980f33223b2cfb8c307fae8ef87d73487d86bf6.
//
// Solidity: event AddFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterAddFactoryBasic(opts *bind.FilterOpts) (*MarketplaceMultiAddFactoryBasicIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "AddFactoryBasic")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiAddFactoryBasicIterator{contract: _MarketplaceMulti.contract, event: "AddFactoryBasic", logs: logs, sub: sub}, nil
}

// WatchAddFactoryBasic is a free log subscription operation binding the contract event 0xdc6ec69d793d0fab8b57c062c980f33223b2cfb8c307fae8ef87d73487d86bf6.
//
// Solidity: event AddFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchAddFactoryBasic(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiAddFactoryBasic) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "AddFactoryBasic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiAddFactoryBasic)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "AddFactoryBasic", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddFactoryBasic is a log parse operation binding the contract event 0xdc6ec69d793d0fab8b57c062c980f33223b2cfb8c307fae8ef87d73487d86bf6.
//
// Solidity: event AddFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseAddFactoryBasic(log types.Log) (*MarketplaceMultiAddFactoryBasic, error) {
	event := new(MarketplaceMultiAddFactoryBasic)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "AddFactoryBasic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiAddFactoryVipIterator is returned from FilterAddFactoryVip and is used to iterate over the raw logs and unpacked data for AddFactoryVip events raised by the MarketplaceMulti contract.
type MarketplaceMultiAddFactoryVipIterator struct {
	Event *MarketplaceMultiAddFactoryVip // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiAddFactoryVipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiAddFactoryVip)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiAddFactoryVip)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiAddFactoryVipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiAddFactoryVipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiAddFactoryVip represents a AddFactoryVip event raised by the MarketplaceMulti contract.
type MarketplaceMultiAddFactoryVip struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddFactoryVip is a free log retrieval operation binding the contract event 0x5b1591e5a9256c7daf633cfb3613b1626430c266dc8c49c7b9c73d31f56b3869.
//
// Solidity: event AddFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterAddFactoryVip(opts *bind.FilterOpts) (*MarketplaceMultiAddFactoryVipIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "AddFactoryVip")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiAddFactoryVipIterator{contract: _MarketplaceMulti.contract, event: "AddFactoryVip", logs: logs, sub: sub}, nil
}

// WatchAddFactoryVip is a free log subscription operation binding the contract event 0x5b1591e5a9256c7daf633cfb3613b1626430c266dc8c49c7b9c73d31f56b3869.
//
// Solidity: event AddFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchAddFactoryVip(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiAddFactoryVip) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "AddFactoryVip")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiAddFactoryVip)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "AddFactoryVip", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddFactoryVip is a log parse operation binding the contract event 0x5b1591e5a9256c7daf633cfb3613b1626430c266dc8c49c7b9c73d31f56b3869.
//
// Solidity: event AddFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseAddFactoryVip(log types.Log) (*MarketplaceMultiAddFactoryVip, error) {
	event := new(MarketplaceMultiAddFactoryVip)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "AddFactoryVip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiBuyMarketItemIterator is returned from FilterBuyMarketItem and is used to iterate over the raw logs and unpacked data for BuyMarketItem events raised by the MarketplaceMulti contract.
type MarketplaceMultiBuyMarketItemIterator struct {
	Event *MarketplaceMultiBuyMarketItem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiBuyMarketItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiBuyMarketItem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiBuyMarketItem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiBuyMarketItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiBuyMarketItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiBuyMarketItem represents a BuyMarketItem event raised by the MarketplaceMulti contract.
type MarketplaceMultiBuyMarketItem struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyMarketItem is a free log retrieval operation binding the contract event 0x8055e776d91c3845104524dc283b0946840289ec1b8ce95ec7d85ed28311287f.
//
// Solidity: event BuyMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterBuyMarketItem(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiBuyMarketItemIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "BuyMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiBuyMarketItemIterator{contract: _MarketplaceMulti.contract, event: "BuyMarketItem", logs: logs, sub: sub}, nil
}

// WatchBuyMarketItem is a free log subscription operation binding the contract event 0x8055e776d91c3845104524dc283b0946840289ec1b8ce95ec7d85ed28311287f.
//
// Solidity: event BuyMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchBuyMarketItem(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiBuyMarketItem, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "BuyMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiBuyMarketItem)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "BuyMarketItem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuyMarketItem is a log parse operation binding the contract event 0x8055e776d91c3845104524dc283b0946840289ec1b8ce95ec7d85ed28311287f.
//
// Solidity: event BuyMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseBuyMarketItem(log types.Log) (*MarketplaceMultiBuyMarketItem, error) {
	event := new(MarketplaceMultiBuyMarketItem)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "BuyMarketItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiCancelSellIterator is returned from FilterCancelSell and is used to iterate over the raw logs and unpacked data for CancelSell events raised by the MarketplaceMulti contract.
type MarketplaceMultiCancelSellIterator struct {
	Event *MarketplaceMultiCancelSell // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiCancelSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiCancelSell)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiCancelSell)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiCancelSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiCancelSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiCancelSell represents a CancelSell event raised by the MarketplaceMulti contract.
type MarketplaceMultiCancelSell struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelSell is a free log retrieval operation binding the contract event 0x29a6902ae24fe755451b145d47596d0840bfbe1c5044ecee155bf4aec677fa02.
//
// Solidity: event CancelSell(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterCancelSell(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiCancelSellIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiCancelSellIterator{contract: _MarketplaceMulti.contract, event: "CancelSell", logs: logs, sub: sub}, nil
}

// WatchCancelSell is a free log subscription operation binding the contract event 0x29a6902ae24fe755451b145d47596d0840bfbe1c5044ecee155bf4aec677fa02.
//
// Solidity: event CancelSell(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchCancelSell(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiCancelSell, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiCancelSell)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "CancelSell", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelSell is a log parse operation binding the contract event 0x29a6902ae24fe755451b145d47596d0840bfbe1c5044ecee155bf4aec677fa02.
//
// Solidity: event CancelSell(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseCancelSell(log types.Log) (*MarketplaceMultiCancelSell, error) {
	event := new(MarketplaceMultiCancelSell)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "CancelSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiCancelSellAuctionIterator is returned from FilterCancelSellAuction and is used to iterate over the raw logs and unpacked data for CancelSellAuction events raised by the MarketplaceMulti contract.
type MarketplaceMultiCancelSellAuctionIterator struct {
	Event *MarketplaceMultiCancelSellAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiCancelSellAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiCancelSellAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiCancelSellAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiCancelSellAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiCancelSellAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiCancelSellAuction represents a CancelSellAuction event raised by the MarketplaceMulti contract.
type MarketplaceMultiCancelSellAuction struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelSellAuction is a free log retrieval operation binding the contract event 0xac465118d818c1752fa3c9dba312f910a9039ae5722ac17b920228ce6306615c.
//
// Solidity: event CancelSellAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterCancelSellAuction(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiCancelSellAuctionIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "CancelSellAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiCancelSellAuctionIterator{contract: _MarketplaceMulti.contract, event: "CancelSellAuction", logs: logs, sub: sub}, nil
}

// WatchCancelSellAuction is a free log subscription operation binding the contract event 0xac465118d818c1752fa3c9dba312f910a9039ae5722ac17b920228ce6306615c.
//
// Solidity: event CancelSellAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchCancelSellAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiCancelSellAuction, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "CancelSellAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiCancelSellAuction)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "CancelSellAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelSellAuction is a log parse operation binding the contract event 0xac465118d818c1752fa3c9dba312f910a9039ae5722ac17b920228ce6306615c.
//
// Solidity: event CancelSellAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseCancelSellAuction(log types.Log) (*MarketplaceMultiCancelSellAuction, error) {
	event := new(MarketplaceMultiCancelSellAuction)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "CancelSellAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiClaimNftIterator is returned from FilterClaimNft and is used to iterate over the raw logs and unpacked data for ClaimNft events raised by the MarketplaceMulti contract.
type MarketplaceMultiClaimNftIterator struct {
	Event *MarketplaceMultiClaimNft // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiClaimNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiClaimNft)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiClaimNft)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiClaimNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiClaimNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiClaimNft represents a ClaimNft event raised by the MarketplaceMulti contract.
type MarketplaceMultiClaimNft struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Buyer       common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimNft is a free log retrieval operation binding the contract event 0xad134c2ad1586b6c1b233ec09b85165f2e35b759c78a884e5f2c8dda9211f2b4.
//
// Solidity: event ClaimNft(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address buyer, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterClaimNft(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiClaimNftIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiClaimNftIterator{contract: _MarketplaceMulti.contract, event: "ClaimNft", logs: logs, sub: sub}, nil
}

// WatchClaimNft is a free log subscription operation binding the contract event 0xad134c2ad1586b6c1b233ec09b85165f2e35b759c78a884e5f2c8dda9211f2b4.
//
// Solidity: event ClaimNft(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address buyer, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchClaimNft(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiClaimNft, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiClaimNft)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "ClaimNft", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimNft is a log parse operation binding the contract event 0xad134c2ad1586b6c1b233ec09b85165f2e35b759c78a884e5f2c8dda9211f2b4.
//
// Solidity: event ClaimNft(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address buyer, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseClaimNft(log types.Log) (*MarketplaceMultiClaimNft, error) {
	event := new(MarketplaceMultiClaimNft)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "ClaimNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiCreateAuctionIterator is returned from FilterCreateAuction and is used to iterate over the raw logs and unpacked data for CreateAuction events raised by the MarketplaceMulti contract.
type MarketplaceMultiCreateAuctionIterator struct {
	Event *MarketplaceMultiCreateAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiCreateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiCreateAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiCreateAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiCreateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiCreateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiCreateAuction represents a CreateAuction event raised by the MarketplaceMulti contract.
type MarketplaceMultiCreateAuction struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	MinBid      *big.Int
	MaxBid      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateAuction is a free log retrieval operation binding the contract event 0x02240a82fcaa45629cce6135ad2760addbffdfb7cfa5590c7a092d3308a97acf.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterCreateAuction(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiCreateAuctionIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiCreateAuctionIterator{contract: _MarketplaceMulti.contract, event: "CreateAuction", logs: logs, sub: sub}, nil
}

// WatchCreateAuction is a free log subscription operation binding the contract event 0x02240a82fcaa45629cce6135ad2760addbffdfb7cfa5590c7a092d3308a97acf.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchCreateAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiCreateAuction, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiCreateAuction)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "CreateAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateAuction is a log parse operation binding the contract event 0x02240a82fcaa45629cce6135ad2760addbffdfb7cfa5590c7a092d3308a97acf.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseCreateAuction(log types.Log) (*MarketplaceMultiCreateAuction, error) {
	event := new(MarketplaceMultiCreateAuction)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "CreateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiCreateMarketItemIterator is returned from FilterCreateMarketItem and is used to iterate over the raw logs and unpacked data for CreateMarketItem events raised by the MarketplaceMulti contract.
type MarketplaceMultiCreateMarketItemIterator struct {
	Event *MarketplaceMultiCreateMarketItem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiCreateMarketItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiCreateMarketItem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiCreateMarketItem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiCreateMarketItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiCreateMarketItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiCreateMarketItem represents a CreateMarketItem event raised by the MarketplaceMulti contract.
type MarketplaceMultiCreateMarketItem struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateMarketItem is a free log retrieval operation binding the contract event 0x666e372f88909b932f48f848ee58673bfa12eecd00c4869a6a8b51bb2dc756b9.
//
// Solidity: event CreateMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price, bool sold)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterCreateMarketItem(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiCreateMarketItemIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "CreateMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiCreateMarketItemIterator{contract: _MarketplaceMulti.contract, event: "CreateMarketItem", logs: logs, sub: sub}, nil
}

// WatchCreateMarketItem is a free log subscription operation binding the contract event 0x666e372f88909b932f48f848ee58673bfa12eecd00c4869a6a8b51bb2dc756b9.
//
// Solidity: event CreateMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price, bool sold)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchCreateMarketItem(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiCreateMarketItem, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "CreateMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiCreateMarketItem)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "CreateMarketItem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateMarketItem is a log parse operation binding the contract event 0x666e372f88909b932f48f848ee58673bfa12eecd00c4869a6a8b51bb2dc756b9.
//
// Solidity: event CreateMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price, bool sold)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseCreateMarketItem(log types.Log) (*MarketplaceMultiCreateMarketItem, error) {
	event := new(MarketplaceMultiCreateMarketItem)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "CreateMarketItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiLoseBidIterator is returned from FilterLoseBid and is used to iterate over the raw logs and unpacked data for LoseBid events raised by the MarketplaceMulti contract.
type MarketplaceMultiLoseBidIterator struct {
	Event *MarketplaceMultiLoseBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiLoseBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiLoseBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiLoseBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiLoseBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiLoseBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiLoseBid represents a LoseBid event raised by the MarketplaceMulti contract.
type MarketplaceMultiLoseBid struct {
	IdOnMarket *big.Int
	User       common.Address
	ReturnBid  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLoseBid is a free log retrieval operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterLoseBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiLoseBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiLoseBidIterator{contract: _MarketplaceMulti.contract, event: "LoseBid", logs: logs, sub: sub}, nil
}

// WatchLoseBid is a free log subscription operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchLoseBid(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiLoseBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiLoseBid)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "LoseBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLoseBid is a log parse operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseLoseBid(log types.Log) (*MarketplaceMultiLoseBid, error) {
	event := new(MarketplaceMultiLoseBid)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "LoseBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewAuctionPackIterator is returned from FilterNewAuctionPack and is used to iterate over the raw logs and unpacked data for NewAuctionPack events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewAuctionPackIterator struct {
	Event *MarketplaceMultiNewAuctionPack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewAuctionPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewAuctionPack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewAuctionPack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewAuctionPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewAuctionPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewAuctionPack represents a NewAuctionPack event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewAuctionPack struct {
	PackId            *big.Int
	StorageIdOnMarket []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewAuctionPack is a free log retrieval operation binding the contract event 0xafdb08065f8d1c171b79c618f0e85ed8641f4067ba8c4cb82583c2aa093bfa57.
//
// Solidity: event NewAuctionPack(uint256 packId, uint256[] storageIdOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewAuctionPack(opts *bind.FilterOpts) (*MarketplaceMultiNewAuctionPackIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewAuctionPack")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewAuctionPackIterator{contract: _MarketplaceMulti.contract, event: "NewAuctionPack", logs: logs, sub: sub}, nil
}

// WatchNewAuctionPack is a free log subscription operation binding the contract event 0xafdb08065f8d1c171b79c618f0e85ed8641f4067ba8c4cb82583c2aa093bfa57.
//
// Solidity: event NewAuctionPack(uint256 packId, uint256[] storageIdOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewAuctionPack(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewAuctionPack) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewAuctionPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewAuctionPack)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewAuctionPack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAuctionPack is a log parse operation binding the contract event 0xafdb08065f8d1c171b79c618f0e85ed8641f4067ba8c4cb82583c2aa093bfa57.
//
// Solidity: event NewAuctionPack(uint256 packId, uint256[] storageIdOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewAuctionPack(log types.Log) (*MarketplaceMultiNewAuctionPack, error) {
	event := new(MarketplaceMultiNewAuctionPack)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewAuctionPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewBidIterator struct {
	Event *MarketplaceMultiNewBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewBid represents a NewBid event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewBid struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Bidder      common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xc96fde502a7126831b4f59269d2f090304bc6b6158e41870ccc833175e9588f8.
//
// Solidity: event NewBid(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address bidder, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceMultiNewBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewBidIterator{contract: _MarketplaceMulti.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xc96fde502a7126831b4f59269d2f090304bc6b6158e41870ccc833175e9588f8.
//
// Solidity: event NewBid(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address bidder, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewBid)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBid is a log parse operation binding the contract event 0xc96fde502a7126831b4f59269d2f090304bc6b6158e41870ccc833175e9588f8.
//
// Solidity: event NewBid(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address bidder, uint256 price)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewBid(log types.Log) (*MarketplaceMultiNewBid, error) {
	event := new(MarketplaceMultiNewBid)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewFeeAddressIterator is returned from FilterNewFeeAddress and is used to iterate over the raw logs and unpacked data for NewFeeAddress events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeAddressIterator struct {
	Event *MarketplaceMultiNewFeeAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewFeeAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewFeeAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewFeeAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewFeeAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewFeeAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewFeeAddress represents a NewFeeAddress event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeAddress struct {
	FeeAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewFeeAddress is a free log retrieval operation binding the contract event 0xdcff558177ef1322e1a9308bd0d2b8b3922d8e05c3ec06323297bdd24db93ef0.
//
// Solidity: event NewFeeAddress(address feeAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewFeeAddress(opts *bind.FilterOpts) (*MarketplaceMultiNewFeeAddressIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewFeeAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewFeeAddressIterator{contract: _MarketplaceMulti.contract, event: "NewFeeAddress", logs: logs, sub: sub}, nil
}

// WatchNewFeeAddress is a free log subscription operation binding the contract event 0xdcff558177ef1322e1a9308bd0d2b8b3922d8e05c3ec06323297bdd24db93ef0.
//
// Solidity: event NewFeeAddress(address feeAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewFeeAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewFeeAddress) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewFeeAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewFeeAddress)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewFeeAddress is a log parse operation binding the contract event 0xdcff558177ef1322e1a9308bd0d2b8b3922d8e05c3ec06323297bdd24db93ef0.
//
// Solidity: event NewFeeAddress(address feeAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewFeeAddress(log types.Log) (*MarketplaceMultiNewFeeAddress, error) {
	event := new(MarketplaceMultiNewFeeAddress)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewFeeCreatorIterator is returned from FilterNewFeeCreator and is used to iterate over the raw logs and unpacked data for NewFeeCreator events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeCreatorIterator struct {
	Event *MarketplaceMultiNewFeeCreator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewFeeCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewFeeCreator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewFeeCreator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewFeeCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewFeeCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewFeeCreator represents a NewFeeCreator event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeCreator struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeCreator is a free log retrieval operation binding the contract event 0xefcc246bd0262a73a8223ffd295abb5ae2cab4c747b7d931061bf4123014e346.
//
// Solidity: event NewFeeCreator(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewFeeCreator(opts *bind.FilterOpts) (*MarketplaceMultiNewFeeCreatorIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewFeeCreator")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewFeeCreatorIterator{contract: _MarketplaceMulti.contract, event: "NewFeeCreator", logs: logs, sub: sub}, nil
}

// WatchNewFeeCreator is a free log subscription operation binding the contract event 0xefcc246bd0262a73a8223ffd295abb5ae2cab4c747b7d931061bf4123014e346.
//
// Solidity: event NewFeeCreator(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewFeeCreator(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewFeeCreator) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewFeeCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewFeeCreator)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeCreator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewFeeCreator is a log parse operation binding the contract event 0xefcc246bd0262a73a8223ffd295abb5ae2cab4c747b7d931061bf4123014e346.
//
// Solidity: event NewFeeCreator(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewFeeCreator(log types.Log) (*MarketplaceMultiNewFeeCreator, error) {
	event := new(MarketplaceMultiNewFeeCreator)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewFeeMarketIterator is returned from FilterNewFeeMarket and is used to iterate over the raw logs and unpacked data for NewFeeMarket events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeMarketIterator struct {
	Event *MarketplaceMultiNewFeeMarket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewFeeMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewFeeMarket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewFeeMarket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewFeeMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewFeeMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewFeeMarket represents a NewFeeMarket event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeMarket struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeMarket is a free log retrieval operation binding the contract event 0xbe3b3f1e97ecc159ccad82985735211f779ad03bdd45040f03fa3fd742deab8a.
//
// Solidity: event NewFeeMarket(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewFeeMarket(opts *bind.FilterOpts) (*MarketplaceMultiNewFeeMarketIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewFeeMarket")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewFeeMarketIterator{contract: _MarketplaceMulti.contract, event: "NewFeeMarket", logs: logs, sub: sub}, nil
}

// WatchNewFeeMarket is a free log subscription operation binding the contract event 0xbe3b3f1e97ecc159ccad82985735211f779ad03bdd45040f03fa3fd742deab8a.
//
// Solidity: event NewFeeMarket(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewFeeMarket(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewFeeMarket) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewFeeMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewFeeMarket)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeMarket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewFeeMarket is a log parse operation binding the contract event 0xbe3b3f1e97ecc159ccad82985735211f779ad03bdd45040f03fa3fd742deab8a.
//
// Solidity: event NewFeeMarket(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewFeeMarket(log types.Log) (*MarketplaceMultiNewFeeMarket, error) {
	event := new(MarketplaceMultiNewFeeMarket)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewFeeRefIterator is returned from FilterNewFeeRef and is used to iterate over the raw logs and unpacked data for NewFeeRef events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeRefIterator struct {
	Event *MarketplaceMultiNewFeeRef // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewFeeRefIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewFeeRef)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewFeeRef)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewFeeRefIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewFeeRefIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewFeeRef represents a NewFeeRef event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewFeeRef struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRef is a free log retrieval operation binding the contract event 0x09c0a82bf115f44d83933f5baafce9d671ef57977cdd2f45487e22fbe5c9b47d.
//
// Solidity: event NewFeeRef(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewFeeRef(opts *bind.FilterOpts) (*MarketplaceMultiNewFeeRefIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewFeeRef")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewFeeRefIterator{contract: _MarketplaceMulti.contract, event: "NewFeeRef", logs: logs, sub: sub}, nil
}

// WatchNewFeeRef is a free log subscription operation binding the contract event 0x09c0a82bf115f44d83933f5baafce9d671ef57977cdd2f45487e22fbe5c9b47d.
//
// Solidity: event NewFeeRef(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewFeeRef(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewFeeRef) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewFeeRef")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewFeeRef)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeRef", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewFeeRef is a log parse operation binding the contract event 0x09c0a82bf115f44d83933f5baafce9d671ef57977cdd2f45487e22fbe5c9b47d.
//
// Solidity: event NewFeeRef(uint256 percent)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewFeeRef(log types.Log) (*MarketplaceMultiNewFeeRef, error) {
	event := new(MarketplaceMultiNewFeeRef)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewFeeRef", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewMainTokenAddressIterator is returned from FilterNewMainTokenAddress and is used to iterate over the raw logs and unpacked data for NewMainTokenAddress events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewMainTokenAddressIterator struct {
	Event *MarketplaceMultiNewMainTokenAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewMainTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewMainTokenAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewMainTokenAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewMainTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewMainTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewMainTokenAddress represents a NewMainTokenAddress event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewMainTokenAddress struct {
	MainToken common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewMainTokenAddress is a free log retrieval operation binding the contract event 0xb560ea6db584a0e0890a534beed59ff9f580da1f25be95f51dd1d580e335e021.
//
// Solidity: event NewMainTokenAddress(address mainToken)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewMainTokenAddress(opts *bind.FilterOpts) (*MarketplaceMultiNewMainTokenAddressIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewMainTokenAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewMainTokenAddressIterator{contract: _MarketplaceMulti.contract, event: "NewMainTokenAddress", logs: logs, sub: sub}, nil
}

// WatchNewMainTokenAddress is a free log subscription operation binding the contract event 0xb560ea6db584a0e0890a534beed59ff9f580da1f25be95f51dd1d580e335e021.
//
// Solidity: event NewMainTokenAddress(address mainToken)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewMainTokenAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewMainTokenAddress) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewMainTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewMainTokenAddress)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewMainTokenAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewMainTokenAddress is a log parse operation binding the contract event 0xb560ea6db584a0e0890a534beed59ff9f580da1f25be95f51dd1d580e335e021.
//
// Solidity: event NewMainTokenAddress(address mainToken)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewMainTokenAddress(log types.Log) (*MarketplaceMultiNewMainTokenAddress, error) {
	event := new(MarketplaceMultiNewMainTokenAddress)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewMainTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewPackIterator is returned from FilterNewPack and is used to iterate over the raw logs and unpacked data for NewPack events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewPackIterator struct {
	Event *MarketplaceMultiNewPack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewPack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewPack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewPack represents a NewPack event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewPack struct {
	Pack        *big.Int
	IdsOnMarket []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPack is a free log retrieval operation binding the contract event 0x46edc46b662430a962a47432ce24ef4cbf0e481e2b27a5f93fec3ea0c52d8887.
//
// Solidity: event NewPack(uint256 pack, uint256[] idsOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewPack(opts *bind.FilterOpts) (*MarketplaceMultiNewPackIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewPack")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewPackIterator{contract: _MarketplaceMulti.contract, event: "NewPack", logs: logs, sub: sub}, nil
}

// WatchNewPack is a free log subscription operation binding the contract event 0x46edc46b662430a962a47432ce24ef4cbf0e481e2b27a5f93fec3ea0c52d8887.
//
// Solidity: event NewPack(uint256 pack, uint256[] idsOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewPack(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewPack) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewPack)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewPack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPack is a log parse operation binding the contract event 0x46edc46b662430a962a47432ce24ef4cbf0e481e2b27a5f93fec3ea0c52d8887.
//
// Solidity: event NewPack(uint256 pack, uint256[] idsOnMarket)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewPack(log types.Log) (*MarketplaceMultiNewPack, error) {
	event := new(MarketplaceMultiNewPack)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiNewRefAddressIterator is returned from FilterNewRefAddress and is used to iterate over the raw logs and unpacked data for NewRefAddress events raised by the MarketplaceMulti contract.
type MarketplaceMultiNewRefAddressIterator struct {
	Event *MarketplaceMultiNewRefAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiNewRefAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiNewRefAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiNewRefAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiNewRefAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiNewRefAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiNewRefAddress represents a NewRefAddress event raised by the MarketplaceMulti contract.
type MarketplaceMultiNewRefAddress struct {
	RefAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRefAddress is a free log retrieval operation binding the contract event 0x7e759132c178bdc3b8f23494e9c61700d192b9fcd42420e58a82811c1e4f8d54.
//
// Solidity: event NewRefAddress(address refAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterNewRefAddress(opts *bind.FilterOpts) (*MarketplaceMultiNewRefAddressIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "NewRefAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiNewRefAddressIterator{contract: _MarketplaceMulti.contract, event: "NewRefAddress", logs: logs, sub: sub}, nil
}

// WatchNewRefAddress is a free log subscription operation binding the contract event 0x7e759132c178bdc3b8f23494e9c61700d192b9fcd42420e58a82811c1e4f8d54.
//
// Solidity: event NewRefAddress(address refAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchNewRefAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiNewRefAddress) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "NewRefAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiNewRefAddress)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "NewRefAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRefAddress is a log parse operation binding the contract event 0x7e759132c178bdc3b8f23494e9c61700d192b9fcd42420e58a82811c1e4f8d54.
//
// Solidity: event NewRefAddress(address refAddress)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseNewRefAddress(log types.Log) (*MarketplaceMultiNewRefAddress, error) {
	event := new(MarketplaceMultiNewRefAddress)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "NewRefAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiRemoveBlackListFeeIterator is returned from FilterRemoveBlackListFee and is used to iterate over the raw logs and unpacked data for RemoveBlackListFee events raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveBlackListFeeIterator struct {
	Event *MarketplaceMultiRemoveBlackListFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiRemoveBlackListFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiRemoveBlackListFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiRemoveBlackListFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiRemoveBlackListFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiRemoveBlackListFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiRemoveBlackListFee represents a RemoveBlackListFee event raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveBlackListFee struct {
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlackListFee is a free log retrieval operation binding the contract event 0x2c3069ddf8f5af7e3a769d0f83893e40875365e3902dab291c5a910b44dced14.
//
// Solidity: event RemoveBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterRemoveBlackListFee(opts *bind.FilterOpts) (*MarketplaceMultiRemoveBlackListFeeIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "RemoveBlackListFee")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiRemoveBlackListFeeIterator{contract: _MarketplaceMulti.contract, event: "RemoveBlackListFee", logs: logs, sub: sub}, nil
}

// WatchRemoveBlackListFee is a free log subscription operation binding the contract event 0x2c3069ddf8f5af7e3a769d0f83893e40875365e3902dab291c5a910b44dced14.
//
// Solidity: event RemoveBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchRemoveBlackListFee(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiRemoveBlackListFee) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "RemoveBlackListFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiRemoveBlackListFee)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveBlackListFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveBlackListFee is a log parse operation binding the contract event 0x2c3069ddf8f5af7e3a769d0f83893e40875365e3902dab291c5a910b44dced14.
//
// Solidity: event RemoveBlackListFee(address[] users)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseRemoveBlackListFee(log types.Log) (*MarketplaceMultiRemoveBlackListFee, error) {
	event := new(MarketplaceMultiRemoveBlackListFee)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveBlackListFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiRemoveFactoryBasicIterator is returned from FilterRemoveFactoryBasic and is used to iterate over the raw logs and unpacked data for RemoveFactoryBasic events raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveFactoryBasicIterator struct {
	Event *MarketplaceMultiRemoveFactoryBasic // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiRemoveFactoryBasicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiRemoveFactoryBasic)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiRemoveFactoryBasic)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiRemoveFactoryBasicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiRemoveFactoryBasicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiRemoveFactoryBasic represents a RemoveFactoryBasic event raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveFactoryBasic struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveFactoryBasic is a free log retrieval operation binding the contract event 0x4f88ae8a1a0a602ddff03d4a85df29951abd94bb171d32b6825fc55bdeecc7bf.
//
// Solidity: event RemoveFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterRemoveFactoryBasic(opts *bind.FilterOpts) (*MarketplaceMultiRemoveFactoryBasicIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "RemoveFactoryBasic")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiRemoveFactoryBasicIterator{contract: _MarketplaceMulti.contract, event: "RemoveFactoryBasic", logs: logs, sub: sub}, nil
}

// WatchRemoveFactoryBasic is a free log subscription operation binding the contract event 0x4f88ae8a1a0a602ddff03d4a85df29951abd94bb171d32b6825fc55bdeecc7bf.
//
// Solidity: event RemoveFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchRemoveFactoryBasic(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiRemoveFactoryBasic) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "RemoveFactoryBasic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiRemoveFactoryBasic)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveFactoryBasic", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveFactoryBasic is a log parse operation binding the contract event 0x4f88ae8a1a0a602ddff03d4a85df29951abd94bb171d32b6825fc55bdeecc7bf.
//
// Solidity: event RemoveFactoryBasic(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseRemoveFactoryBasic(log types.Log) (*MarketplaceMultiRemoveFactoryBasic, error) {
	event := new(MarketplaceMultiRemoveFactoryBasic)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveFactoryBasic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMultiRemoveFactoryVipIterator is returned from FilterRemoveFactoryVip and is used to iterate over the raw logs and unpacked data for RemoveFactoryVip events raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveFactoryVipIterator struct {
	Event *MarketplaceMultiRemoveFactoryVip // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceMultiRemoveFactoryVipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMultiRemoveFactoryVip)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceMultiRemoveFactoryVip)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceMultiRemoveFactoryVipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMultiRemoveFactoryVipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMultiRemoveFactoryVip represents a RemoveFactoryVip event raised by the MarketplaceMulti contract.
type MarketplaceMultiRemoveFactoryVip struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveFactoryVip is a free log retrieval operation binding the contract event 0xef7affc8c54e6b9a0413bdc46ec5c49dd3ecd035a5e6763f8c0fac4cd256ff64.
//
// Solidity: event RemoveFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) FilterRemoveFactoryVip(opts *bind.FilterOpts) (*MarketplaceMultiRemoveFactoryVipIterator, error) {

	logs, sub, err := _MarketplaceMulti.contract.FilterLogs(opts, "RemoveFactoryVip")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMultiRemoveFactoryVipIterator{contract: _MarketplaceMulti.contract, event: "RemoveFactoryVip", logs: logs, sub: sub}, nil
}

// WatchRemoveFactoryVip is a free log subscription operation binding the contract event 0xef7affc8c54e6b9a0413bdc46ec5c49dd3ecd035a5e6763f8c0fac4cd256ff64.
//
// Solidity: event RemoveFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) WatchRemoveFactoryVip(opts *bind.WatchOpts, sink chan<- *MarketplaceMultiRemoveFactoryVip) (event.Subscription, error) {

	logs, sub, err := _MarketplaceMulti.contract.WatchLogs(opts, "RemoveFactoryVip")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMultiRemoveFactoryVip)
				if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveFactoryVip", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveFactoryVip is a log parse operation binding the contract event 0xef7affc8c54e6b9a0413bdc46ec5c49dd3ecd035a5e6763f8c0fac4cd256ff64.
//
// Solidity: event RemoveFactoryVip(address factory)
func (_MarketplaceMulti *MarketplaceMultiFilterer) ParseRemoveFactoryVip(log types.Log) (*MarketplaceMultiRemoveFactoryVip, error) {
	event := new(MarketplaceMultiRemoveFactoryVip)
	if err := _MarketplaceMulti.contract.UnpackLog(event, "RemoveFactoryVip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
