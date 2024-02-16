// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CancelSellAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ClaimNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"}],\"name\":\"CreateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnBid\",\"type\":\"uint256\"}],\"name\":\"LoseBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"storageIdOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"NewAuctionPack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelSellAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchClaimNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"batchSaleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"cancelSellAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"claimNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getListingItemInOnePackAuction\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"saleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyMarketItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CancelSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"CreateMarketItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"idsOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"NewPack\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceItem\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"}],\"name\":\"Sale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idsOnMarket\",\"type\":\"uint256[]\"},{\"internalType\":\"enumIMarketplaceFeature.BATCH_BUY_TYPE\",\"name\":\"buyType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"batchBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_idOnMarket\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceItem\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"}],\"name\":\"batchSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"cancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIdOnMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"}],\"internalType\":\"structLibMarketplaceStorage.MarketItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getIsPackSold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pack\",\"type\":\"uint256\"}],\"name\":\"getListingItemInOnePack\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pack\",\"type\":\"uint256\"}],\"name\":\"getPack\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPackSold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fromId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"}],\"internalType\":\"structLibMarketplaceStorage.PackData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"AddBlackListFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"AddFactoryBasic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"AddFactoryVip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"NewFeeAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"NewFeeRef\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mainToken\",\"type\":\"address\"}],\"name\":\"NewMainTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refAddress\",\"type\":\"address\"}],\"name\":\"NewRefAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"RemoveBlackListFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RemoveFactoryBasic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RemoveFactoryVip\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkBlackListFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"checkFactoryBasic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"checkFactoryVip\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getBuyData\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeCreator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRef\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRefAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"removeBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"setBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeRef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MainToken\",\"type\":\"address\"}],\"name\":\"setMainTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_RefAddress\",\"type\":\"address\"}],\"name\":\"setRefAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_Marketplace *MarketplaceCaller) CheckBlackListFee(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "checkBlackListFee", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_Marketplace *MarketplaceSession) CheckBlackListFee(user common.Address) (bool, error) {
	return _Marketplace.Contract.CheckBlackListFee(&_Marketplace.CallOpts, user)
}

// CheckBlackListFee is a free data retrieval call binding the contract method 0x8c874159.
//
// Solidity: function checkBlackListFee(address user) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) CheckBlackListFee(user common.Address) (bool, error) {
	return _Marketplace.Contract.CheckBlackListFee(&_Marketplace.CallOpts, user)
}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_Marketplace *MarketplaceCaller) CheckFactoryBasic(opts *bind.CallOpts, proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "checkFactoryBasic", proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_Marketplace *MarketplaceSession) CheckFactoryBasic(proxy common.Address) (bool, error) {
	return _Marketplace.Contract.CheckFactoryBasic(&_Marketplace.CallOpts, proxy)
}

// CheckFactoryBasic is a free data retrieval call binding the contract method 0x9f02454c.
//
// Solidity: function checkFactoryBasic(address proxy) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) CheckFactoryBasic(proxy common.Address) (bool, error) {
	return _Marketplace.Contract.CheckFactoryBasic(&_Marketplace.CallOpts, proxy)
}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_Marketplace *MarketplaceCaller) CheckFactoryVip(opts *bind.CallOpts, proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "checkFactoryVip", proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_Marketplace *MarketplaceSession) CheckFactoryVip(proxy common.Address) (bool, error) {
	return _Marketplace.Contract.CheckFactoryVip(&_Marketplace.CallOpts, proxy)
}

// CheckFactoryVip is a free data retrieval call binding the contract method 0xc6abbd7e.
//
// Solidity: function checkFactoryVip(address proxy) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) CheckFactoryVip(proxy common.Address) (bool, error) {
	return _Marketplace.Contract.CheckFactoryVip(&_Marketplace.CallOpts, proxy)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_Marketplace *MarketplaceCaller) GetBuyData(opts *bind.CallOpts, price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBuyData", price_, factory_, seller_, tokenId_)

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
func (_Marketplace *MarketplaceSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _Marketplace.Contract.GetBuyData(&_Marketplace.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_Marketplace *MarketplaceCallerSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _Marketplace.Contract.GetBuyData(&_Marketplace.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Marketplace *MarketplaceCaller) GetFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Marketplace *MarketplaceSession) GetFeeAddress() (common.Address, error) {
	return _Marketplace.Contract.GetFeeAddress(&_Marketplace.CallOpts)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetFeeAddress() (common.Address, error) {
	return _Marketplace.Contract.GetFeeAddress(&_Marketplace.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetFeeCreator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getFeeCreator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetFeeCreator() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeCreator(&_Marketplace.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetFeeCreator() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeCreator(&_Marketplace.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetFeeMarket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getFeeMarket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetFeeMarket() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeMarket(&_Marketplace.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetFeeMarket() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeMarket(&_Marketplace.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetFeeRef(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getFeeRef")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetFeeRef() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeRef(&_Marketplace.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetFeeRef() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeRef(&_Marketplace.CallOpts)
}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_Marketplace *MarketplaceCaller) GetIdOnMarket(opts *bind.CallOpts, id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getIdOnMarket", id)

	if err != nil {
		return *new(LibMarketplaceStorageMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new(LibMarketplaceStorageMarketItem)).(*LibMarketplaceStorageMarketItem)

	return out0, err

}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_Marketplace *MarketplaceSession) GetIdOnMarket(id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	return _Marketplace.Contract.GetIdOnMarket(&_Marketplace.CallOpts, id)
}

// GetIdOnMarket is a free data retrieval call binding the contract method 0xd8b0a70d.
//
// Solidity: function getIdOnMarket(uint256 id) view returns((uint256,uint256,address,address,address,address,uint256,bool,uint256))
func (_Marketplace *MarketplaceCallerSession) GetIdOnMarket(id *big.Int) (LibMarketplaceStorageMarketItem, error) {
	return _Marketplace.Contract.GetIdOnMarket(&_Marketplace.CallOpts, id)
}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_Marketplace *MarketplaceCaller) GetIsPackSold(opts *bind.CallOpts, _pack *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getIsPackSold", _pack)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_Marketplace *MarketplaceSession) GetIsPackSold(_pack *big.Int) (bool, error) {
	return _Marketplace.Contract.GetIsPackSold(&_Marketplace.CallOpts, _pack)
}

// GetIsPackSold is a free data retrieval call binding the contract method 0xb35b6f71.
//
// Solidity: function getIsPackSold(uint256 _pack) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) GetIsPackSold(_pack *big.Int) (bool, error) {
	return _Marketplace.Contract.GetIsPackSold(&_Marketplace.CallOpts, _pack)
}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetListingItemInOnePack(opts *bind.CallOpts, _pack *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListingItemInOnePack", _pack)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetListingItemInOnePack(_pack *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetListingItemInOnePack(&_Marketplace.CallOpts, _pack)
}

// GetListingItemInOnePack is a free data retrieval call binding the contract method 0x065a4c52.
//
// Solidity: function getListingItemInOnePack(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetListingItemInOnePack(_pack *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetListingItemInOnePack(&_Marketplace.CallOpts, _pack)
}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetListingItemInOnePackAuction(opts *bind.CallOpts, _pack *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListingItemInOnePackAuction", _pack)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetListingItemInOnePackAuction(_pack *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetListingItemInOnePackAuction(&_Marketplace.CallOpts, _pack)
}

// GetListingItemInOnePackAuction is a free data retrieval call binding the contract method 0x4969699d.
//
// Solidity: function getListingItemInOnePackAuction(uint256 _pack) view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetListingItemInOnePackAuction(_pack *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetListingItemInOnePackAuction(&_Marketplace.CallOpts, _pack)
}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_Marketplace *MarketplaceCaller) GetMainTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getMainTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_Marketplace *MarketplaceSession) GetMainTokenAddress() (common.Address, error) {
	return _Marketplace.Contract.GetMainTokenAddress(&_Marketplace.CallOpts)
}

// GetMainTokenAddress is a free data retrieval call binding the contract method 0xd3b61a59.
//
// Solidity: function getMainTokenAddress() view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetMainTokenAddress() (common.Address, error) {
	return _Marketplace.Contract.GetMainTokenAddress(&_Marketplace.CallOpts)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_Marketplace *MarketplaceCaller) GetPack(opts *bind.CallOpts, pack *big.Int) (LibMarketplaceStoragePackData, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getPack", pack)

	if err != nil {
		return *new(LibMarketplaceStoragePackData), err
	}

	out0 := *abi.ConvertType(out[0], new(LibMarketplaceStoragePackData)).(*LibMarketplaceStoragePackData)

	return out0, err

}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_Marketplace *MarketplaceSession) GetPack(pack *big.Int) (LibMarketplaceStoragePackData, error) {
	return _Marketplace.Contract.GetPack(&_Marketplace.CallOpts, pack)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 pack) view returns((bool,uint256,uint256))
func (_Marketplace *MarketplaceCallerSession) GetPack(pack *big.Int) (LibMarketplaceStoragePackData, error) {
	return _Marketplace.Contract.GetPack(&_Marketplace.CallOpts, pack)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Marketplace *MarketplaceCaller) GetRefAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRefAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Marketplace *MarketplaceSession) GetRefAddress() (common.Address, error) {
	return _Marketplace.Contract.GetRefAddress(&_Marketplace.CallOpts)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetRefAddress() (common.Address, error) {
	return _Marketplace.Contract.GetRefAddress(&_Marketplace.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactor) Bid(opts *bind.TransactOpts, _idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "Bid", _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Marketplace *MarketplaceSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Bid(&_Marketplace.TransactOpts, _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Bid(&_Marketplace.TransactOpts, _idOnMarket, _price)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) Buy(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "Buy", _idOnMarket)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceSession) Buy(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _idOnMarket)
}

// Buy is a paid mutator transaction binding the contract method 0x3e328218.
//
// Solidity: function Buy(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) Buy(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _idOnMarket)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceTransactor) Sale(opts *bind.TransactOpts, _tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "Sale", _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceSession) Sale(_tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Sale(&_Marketplace.TransactOpts, _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// Sale is a paid mutator transaction binding the contract method 0x32f7b044.
//
// Solidity: function Sale(uint256 _tokenId, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceTransactorSession) Sale(_tokenId *big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Sale(&_Marketplace.TransactOpts, _tokenId, _factoryAddress, _priceItem, _contractHold)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceTransactor) AddFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addFactoryBasic", proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AddFactoryBasic(&_Marketplace.TransactOpts, proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceTransactorSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AddFactoryBasic(&_Marketplace.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceTransactor) AddFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addFactoryVip", proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AddFactoryVip(&_Marketplace.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceTransactorSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AddFactoryVip(&_Marketplace.TransactOpts, proxy)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactor) BatchBuy(opts *bind.TransactOpts, packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchBuy", packId, _idsOnMarket, buyType, _price)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_Marketplace *MarketplaceSession) BatchBuy(packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchBuy(&_Marketplace.TransactOpts, packId, _idsOnMarket, buyType, _price)
}

// BatchBuy is a paid mutator transaction binding the contract method 0xfab74e2f.
//
// Solidity: function batchBuy(uint256 packId, uint256[] _idsOnMarket, uint8 buyType, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchBuy(packId *big.Int, _idsOnMarket []*big.Int, buyType uint8, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchBuy(&_Marketplace.TransactOpts, packId, _idsOnMarket, buyType, _price)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) BatchCancelSell(opts *bind.TransactOpts, packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchCancelSell", packId, _idOnMarket)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_Marketplace *MarketplaceSession) BatchCancelSell(packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchCancelSell(&_Marketplace.TransactOpts, packId, _idOnMarket)
}

// BatchCancelSell is a paid mutator transaction binding the contract method 0x9cd3ce53.
//
// Solidity: function batchCancelSell(uint256 packId, uint256[] _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchCancelSell(packId *big.Int, _idOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchCancelSell(&_Marketplace.TransactOpts, packId, _idOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) BatchCancelSellAuction(opts *bind.TransactOpts, packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchCancelSellAuction", packId, idsOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_Marketplace *MarketplaceSession) BatchCancelSellAuction(packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchCancelSellAuction(&_Marketplace.TransactOpts, packId, idsOnMarket)
}

// BatchCancelSellAuction is a paid mutator transaction binding the contract method 0xbc855c61.
//
// Solidity: function batchCancelSellAuction(uint256 packId, uint256[] idsOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchCancelSellAuction(packId *big.Int, idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchCancelSellAuction(&_Marketplace.TransactOpts, packId, idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) BatchClaimNft(opts *bind.TransactOpts, pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchClaimNft", pack, _idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_Marketplace *MarketplaceSession) BatchClaimNft(pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchClaimNft(&_Marketplace.TransactOpts, pack, _idsOnMarket)
}

// BatchClaimNft is a paid mutator transaction binding the contract method 0xd37d5ba7.
//
// Solidity: function batchClaimNft(uint256 pack, uint256[] _idsOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchClaimNft(pack *big.Int, _idsOnMarket []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchClaimNft(&_Marketplace.TransactOpts, pack, _idsOnMarket)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceTransactor) BatchSale(opts *bind.TransactOpts, _tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchSale", _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceSession) BatchSale(_tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchSale(&_Marketplace.TransactOpts, _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSale is a paid mutator transaction binding the contract method 0xc35c84c1.
//
// Solidity: function batchSale(uint256[] _tokenIds, address _factoryAddress, uint256 _priceItem, address _contractHold) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchSale(_tokenIds []*big.Int, _factoryAddress common.Address, _priceItem *big.Int, _contractHold common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchSale(&_Marketplace.TransactOpts, _tokenIds, _factoryAddress, _priceItem, _contractHold)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceTransactor) BatchSaleAuction(opts *bind.TransactOpts, _tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "batchSaleAuction", _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceSession) BatchSaleAuction(_tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchSaleAuction(&_Marketplace.TransactOpts, _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// BatchSaleAuction is a paid mutator transaction binding the contract method 0xf5bfd8b7.
//
// Solidity: function batchSaleAuction(uint256[] _tokenIds, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceTransactorSession) BatchSaleAuction(_tokenIds []*big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BatchSaleAuction(&_Marketplace.TransactOpts, _tokenIds, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) CancelSell(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelSell", _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSell(&_Marketplace.TransactOpts, _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSell(&_Marketplace.TransactOpts, _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) CancelSellAuction(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelSellAuction", _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceSession) CancelSellAuction(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSellAuction(&_Marketplace.TransactOpts, _idOnMarket)
}

// CancelSellAuction is a paid mutator transaction binding the contract method 0xcef93338.
//
// Solidity: function cancelSellAuction(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelSellAuction(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSellAuction(&_Marketplace.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactor) ClaimNft(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "claimNft", _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ClaimNft(&_Marketplace.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Marketplace *MarketplaceTransactorSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ClaimNft(&_Marketplace.TransactOpts, _idOnMarket)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceSession) Migrate() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate(&_Marketplace.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactorSession) Migrate() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate(&_Marketplace.TransactOpts)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactor) Migrate0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "migrate0")
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceSession) Migrate0() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate0(&_Marketplace.TransactOpts)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactorSession) Migrate0() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate0(&_Marketplace.TransactOpts)
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactor) Migrate1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "migrate1")
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceSession) Migrate1() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate1(&_Marketplace.TransactOpts)
}

// Migrate1 is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_Marketplace *MarketplaceTransactorSession) Migrate1() (*types.Transaction, error) {
	return _Marketplace.Contract.Migrate1(&_Marketplace.TransactOpts)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceTransactor) RemoveBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "removeBlackListFee", user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveBlackListFee(&_Marketplace.TransactOpts, user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceTransactorSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveBlackListFee(&_Marketplace.TransactOpts, user)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceTransactor) RemoveFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "removeFactoryBasic", proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFactoryBasic(&_Marketplace.TransactOpts, proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Marketplace *MarketplaceTransactorSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFactoryBasic(&_Marketplace.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceTransactor) RemoveFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "removeFactoryVip", proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFactoryVip(&_Marketplace.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Marketplace *MarketplaceTransactorSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RemoveFactoryVip(&_Marketplace.TransactOpts, proxy)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceTransactor) SaleAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "saleAuction", _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SaleAuction(&_Marketplace.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x199289b9.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid, address _contractHold, address _factory) returns()
func (_Marketplace *MarketplaceTransactorSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _maxBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SaleAuction(&_Marketplace.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _maxBid, _contractHold, _factory)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceTransactor) SetBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setBlackListFee", user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBlackListFee(&_Marketplace.TransactOpts, user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Marketplace *MarketplaceTransactorSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBlackListFee(&_Marketplace.TransactOpts, user)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Marketplace *MarketplaceTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Marketplace *MarketplaceSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeAddress(&_Marketplace.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeAddress(&_Marketplace.TransactOpts, _feeAddress)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactor) SetFeeCreator(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setFeeCreator", percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Marketplace *MarketplaceSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeCreator(&_Marketplace.TransactOpts, percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactorSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeCreator(&_Marketplace.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactor) SetFeeMarket(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setFeeMarket", percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Marketplace *MarketplaceSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeMarket(&_Marketplace.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactorSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeMarket(&_Marketplace.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactor) SetFeeRef(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setFeeRef", percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Marketplace *MarketplaceSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeRef(&_Marketplace.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Marketplace *MarketplaceTransactorSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeRef(&_Marketplace.TransactOpts, percent)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Marketplace *MarketplaceTransactor) SetMainTokenAddress(opts *bind.TransactOpts, _MainToken common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMainTokenAddress", _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Marketplace *MarketplaceSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMainTokenAddress(&_Marketplace.TransactOpts, _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMainTokenAddress(&_Marketplace.TransactOpts, _MainToken)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Marketplace *MarketplaceTransactor) SetRefAddress(opts *bind.TransactOpts, _RefAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRefAddress", _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Marketplace *MarketplaceSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRefAddress(&_Marketplace.TransactOpts, _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRefAddress(&_Marketplace.TransactOpts, _RefAddress)
}

// MarketplaceAddBlackListFeeIterator is returned from FilterAddBlackListFee and is used to iterate over the raw logs and unpacked data for AddBlackListFee events raised by the Marketplace contract.
type MarketplaceAddBlackListFeeIterator struct {
	Event *MarketplaceAddBlackListFee // Event containing the contract specifics and raw log

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
func (it *MarketplaceAddBlackListFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAddBlackListFee)
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
		it.Event = new(MarketplaceAddBlackListFee)
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
func (it *MarketplaceAddBlackListFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAddBlackListFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAddBlackListFee represents a AddBlackListFee event raised by the Marketplace contract.
type MarketplaceAddBlackListFee struct {
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddBlackListFee is a free log retrieval operation binding the contract event 0x6c23ff28a8bb18a73e79a449be6a4435f5123b723715d33899f0b39e2b2beb07.
//
// Solidity: event AddBlackListFee(address[] users)
func (_Marketplace *MarketplaceFilterer) FilterAddBlackListFee(opts *bind.FilterOpts) (*MarketplaceAddBlackListFeeIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AddBlackListFee")
	if err != nil {
		return nil, err
	}
	return &MarketplaceAddBlackListFeeIterator{contract: _Marketplace.contract, event: "AddBlackListFee", logs: logs, sub: sub}, nil
}

// WatchAddBlackListFee is a free log subscription operation binding the contract event 0x6c23ff28a8bb18a73e79a449be6a4435f5123b723715d33899f0b39e2b2beb07.
//
// Solidity: event AddBlackListFee(address[] users)
func (_Marketplace *MarketplaceFilterer) WatchAddBlackListFee(opts *bind.WatchOpts, sink chan<- *MarketplaceAddBlackListFee) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AddBlackListFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAddBlackListFee)
				if err := _Marketplace.contract.UnpackLog(event, "AddBlackListFee", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseAddBlackListFee(log types.Log) (*MarketplaceAddBlackListFee, error) {
	event := new(MarketplaceAddBlackListFee)
	if err := _Marketplace.contract.UnpackLog(event, "AddBlackListFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceAddFactoryBasicIterator is returned from FilterAddFactoryBasic and is used to iterate over the raw logs and unpacked data for AddFactoryBasic events raised by the Marketplace contract.
type MarketplaceAddFactoryBasicIterator struct {
	Event *MarketplaceAddFactoryBasic // Event containing the contract specifics and raw log

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
func (it *MarketplaceAddFactoryBasicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAddFactoryBasic)
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
		it.Event = new(MarketplaceAddFactoryBasic)
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
func (it *MarketplaceAddFactoryBasicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAddFactoryBasicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAddFactoryBasic represents a AddFactoryBasic event raised by the Marketplace contract.
type MarketplaceAddFactoryBasic struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddFactoryBasic is a free log retrieval operation binding the contract event 0xdc6ec69d793d0fab8b57c062c980f33223b2cfb8c307fae8ef87d73487d86bf6.
//
// Solidity: event AddFactoryBasic(address factory)
func (_Marketplace *MarketplaceFilterer) FilterAddFactoryBasic(opts *bind.FilterOpts) (*MarketplaceAddFactoryBasicIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AddFactoryBasic")
	if err != nil {
		return nil, err
	}
	return &MarketplaceAddFactoryBasicIterator{contract: _Marketplace.contract, event: "AddFactoryBasic", logs: logs, sub: sub}, nil
}

// WatchAddFactoryBasic is a free log subscription operation binding the contract event 0xdc6ec69d793d0fab8b57c062c980f33223b2cfb8c307fae8ef87d73487d86bf6.
//
// Solidity: event AddFactoryBasic(address factory)
func (_Marketplace *MarketplaceFilterer) WatchAddFactoryBasic(opts *bind.WatchOpts, sink chan<- *MarketplaceAddFactoryBasic) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AddFactoryBasic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAddFactoryBasic)
				if err := _Marketplace.contract.UnpackLog(event, "AddFactoryBasic", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseAddFactoryBasic(log types.Log) (*MarketplaceAddFactoryBasic, error) {
	event := new(MarketplaceAddFactoryBasic)
	if err := _Marketplace.contract.UnpackLog(event, "AddFactoryBasic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceAddFactoryVipIterator is returned from FilterAddFactoryVip and is used to iterate over the raw logs and unpacked data for AddFactoryVip events raised by the Marketplace contract.
type MarketplaceAddFactoryVipIterator struct {
	Event *MarketplaceAddFactoryVip // Event containing the contract specifics and raw log

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
func (it *MarketplaceAddFactoryVipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAddFactoryVip)
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
		it.Event = new(MarketplaceAddFactoryVip)
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
func (it *MarketplaceAddFactoryVipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAddFactoryVipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAddFactoryVip represents a AddFactoryVip event raised by the Marketplace contract.
type MarketplaceAddFactoryVip struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddFactoryVip is a free log retrieval operation binding the contract event 0x5b1591e5a9256c7daf633cfb3613b1626430c266dc8c49c7b9c73d31f56b3869.
//
// Solidity: event AddFactoryVip(address factory)
func (_Marketplace *MarketplaceFilterer) FilterAddFactoryVip(opts *bind.FilterOpts) (*MarketplaceAddFactoryVipIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AddFactoryVip")
	if err != nil {
		return nil, err
	}
	return &MarketplaceAddFactoryVipIterator{contract: _Marketplace.contract, event: "AddFactoryVip", logs: logs, sub: sub}, nil
}

// WatchAddFactoryVip is a free log subscription operation binding the contract event 0x5b1591e5a9256c7daf633cfb3613b1626430c266dc8c49c7b9c73d31f56b3869.
//
// Solidity: event AddFactoryVip(address factory)
func (_Marketplace *MarketplaceFilterer) WatchAddFactoryVip(opts *bind.WatchOpts, sink chan<- *MarketplaceAddFactoryVip) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AddFactoryVip")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAddFactoryVip)
				if err := _Marketplace.contract.UnpackLog(event, "AddFactoryVip", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseAddFactoryVip(log types.Log) (*MarketplaceAddFactoryVip, error) {
	event := new(MarketplaceAddFactoryVip)
	if err := _Marketplace.contract.UnpackLog(event, "AddFactoryVip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceBuyMarketItemIterator is returned from FilterBuyMarketItem and is used to iterate over the raw logs and unpacked data for BuyMarketItem events raised by the Marketplace contract.
type MarketplaceBuyMarketItemIterator struct {
	Event *MarketplaceBuyMarketItem // Event containing the contract specifics and raw log

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
func (it *MarketplaceBuyMarketItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBuyMarketItem)
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
		it.Event = new(MarketplaceBuyMarketItem)
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
func (it *MarketplaceBuyMarketItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBuyMarketItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBuyMarketItem represents a BuyMarketItem event raised by the Marketplace contract.
type MarketplaceBuyMarketItem struct {
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
func (_Marketplace *MarketplaceFilterer) FilterBuyMarketItem(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceBuyMarketItemIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BuyMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBuyMarketItemIterator{contract: _Marketplace.contract, event: "BuyMarketItem", logs: logs, sub: sub}, nil
}

// WatchBuyMarketItem is a free log subscription operation binding the contract event 0x8055e776d91c3845104524dc283b0946840289ec1b8ce95ec7d85ed28311287f.
//
// Solidity: event BuyMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchBuyMarketItem(opts *bind.WatchOpts, sink chan<- *MarketplaceBuyMarketItem, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BuyMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBuyMarketItem)
				if err := _Marketplace.contract.UnpackLog(event, "BuyMarketItem", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseBuyMarketItem(log types.Log) (*MarketplaceBuyMarketItem, error) {
	event := new(MarketplaceBuyMarketItem)
	if err := _Marketplace.contract.UnpackLog(event, "BuyMarketItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCancelSellIterator is returned from FilterCancelSell and is used to iterate over the raw logs and unpacked data for CancelSell events raised by the Marketplace contract.
type MarketplaceCancelSellIterator struct {
	Event *MarketplaceCancelSell // Event containing the contract specifics and raw log

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
func (it *MarketplaceCancelSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCancelSell)
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
		it.Event = new(MarketplaceCancelSell)
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
func (it *MarketplaceCancelSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCancelSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCancelSell represents a CancelSell event raised by the Marketplace contract.
type MarketplaceCancelSell struct {
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
func (_Marketplace *MarketplaceFilterer) FilterCancelSell(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceCancelSellIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCancelSellIterator{contract: _Marketplace.contract, event: "CancelSell", logs: logs, sub: sub}, nil
}

// WatchCancelSell is a free log subscription operation binding the contract event 0x29a6902ae24fe755451b145d47596d0840bfbe1c5044ecee155bf4aec677fa02.
//
// Solidity: event CancelSell(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_Marketplace *MarketplaceFilterer) WatchCancelSell(opts *bind.WatchOpts, sink chan<- *MarketplaceCancelSell, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCancelSell)
				if err := _Marketplace.contract.UnpackLog(event, "CancelSell", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseCancelSell(log types.Log) (*MarketplaceCancelSell, error) {
	event := new(MarketplaceCancelSell)
	if err := _Marketplace.contract.UnpackLog(event, "CancelSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCancelSellAuctionIterator is returned from FilterCancelSellAuction and is used to iterate over the raw logs and unpacked data for CancelSellAuction events raised by the Marketplace contract.
type MarketplaceCancelSellAuctionIterator struct {
	Event *MarketplaceCancelSellAuction // Event containing the contract specifics and raw log

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
func (it *MarketplaceCancelSellAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCancelSellAuction)
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
		it.Event = new(MarketplaceCancelSellAuction)
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
func (it *MarketplaceCancelSellAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCancelSellAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCancelSellAuction represents a CancelSellAuction event raised by the Marketplace contract.
type MarketplaceCancelSellAuction struct {
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
func (_Marketplace *MarketplaceFilterer) FilterCancelSellAuction(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceCancelSellAuctionIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CancelSellAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCancelSellAuctionIterator{contract: _Marketplace.contract, event: "CancelSellAuction", logs: logs, sub: sub}, nil
}

// WatchCancelSellAuction is a free log subscription operation binding the contract event 0xac465118d818c1752fa3c9dba312f910a9039ae5722ac17b920228ce6306615c.
//
// Solidity: event CancelSellAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_Marketplace *MarketplaceFilterer) WatchCancelSellAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceCancelSellAuction, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CancelSellAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCancelSellAuction)
				if err := _Marketplace.contract.UnpackLog(event, "CancelSellAuction", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseCancelSellAuction(log types.Log) (*MarketplaceCancelSellAuction, error) {
	event := new(MarketplaceCancelSellAuction)
	if err := _Marketplace.contract.UnpackLog(event, "CancelSellAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceClaimNftIterator is returned from FilterClaimNft and is used to iterate over the raw logs and unpacked data for ClaimNft events raised by the Marketplace contract.
type MarketplaceClaimNftIterator struct {
	Event *MarketplaceClaimNft // Event containing the contract specifics and raw log

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
func (it *MarketplaceClaimNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceClaimNft)
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
		it.Event = new(MarketplaceClaimNft)
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
func (it *MarketplaceClaimNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceClaimNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceClaimNft represents a ClaimNft event raised by the Marketplace contract.
type MarketplaceClaimNft struct {
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
func (_Marketplace *MarketplaceFilterer) FilterClaimNft(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceClaimNftIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceClaimNftIterator{contract: _Marketplace.contract, event: "ClaimNft", logs: logs, sub: sub}, nil
}

// WatchClaimNft is a free log subscription operation binding the contract event 0xad134c2ad1586b6c1b233ec09b85165f2e35b759c78a884e5f2c8dda9211f2b4.
//
// Solidity: event ClaimNft(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchClaimNft(opts *bind.WatchOpts, sink chan<- *MarketplaceClaimNft, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceClaimNft)
				if err := _Marketplace.contract.UnpackLog(event, "ClaimNft", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseClaimNft(log types.Log) (*MarketplaceClaimNft, error) {
	event := new(MarketplaceClaimNft)
	if err := _Marketplace.contract.UnpackLog(event, "ClaimNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateAuctionIterator is returned from FilterCreateAuction and is used to iterate over the raw logs and unpacked data for CreateAuction events raised by the Marketplace contract.
type MarketplaceCreateAuctionIterator struct {
	Event *MarketplaceCreateAuction // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateAuction)
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
		it.Event = new(MarketplaceCreateAuction)
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
func (it *MarketplaceCreateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateAuction represents a CreateAuction event raised by the Marketplace contract.
type MarketplaceCreateAuction struct {
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
func (_Marketplace *MarketplaceFilterer) FilterCreateAuction(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceCreateAuctionIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateAuctionIterator{contract: _Marketplace.contract, event: "CreateAuction", logs: logs, sub: sub}, nil
}

// WatchCreateAuction is a free log subscription operation binding the contract event 0x02240a82fcaa45629cce6135ad2760addbffdfb7cfa5590c7a092d3308a97acf.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid, uint256 _maxBid)
func (_Marketplace *MarketplaceFilterer) WatchCreateAuction(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateAuction, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateAuction)
				if err := _Marketplace.contract.UnpackLog(event, "CreateAuction", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseCreateAuction(log types.Log) (*MarketplaceCreateAuction, error) {
	event := new(MarketplaceCreateAuction)
	if err := _Marketplace.contract.UnpackLog(event, "CreateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateMarketItemIterator is returned from FilterCreateMarketItem and is used to iterate over the raw logs and unpacked data for CreateMarketItem events raised by the Marketplace contract.
type MarketplaceCreateMarketItemIterator struct {
	Event *MarketplaceCreateMarketItem // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateMarketItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateMarketItem)
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
		it.Event = new(MarketplaceCreateMarketItem)
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
func (it *MarketplaceCreateMarketItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateMarketItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateMarketItem represents a CreateMarketItem event raised by the Marketplace contract.
type MarketplaceCreateMarketItem struct {
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
func (_Marketplace *MarketplaceFilterer) FilterCreateMarketItem(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceCreateMarketItemIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateMarketItemIterator{contract: _Marketplace.contract, event: "CreateMarketItem", logs: logs, sub: sub}, nil
}

// WatchCreateMarketItem is a free log subscription operation binding the contract event 0x666e372f88909b932f48f848ee58673bfa12eecd00c4869a6a8b51bb2dc756b9.
//
// Solidity: event CreateMarketItem(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner, uint256 price, bool sold)
func (_Marketplace *MarketplaceFilterer) WatchCreateMarketItem(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateMarketItem, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateMarketItem", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateMarketItem)
				if err := _Marketplace.contract.UnpackLog(event, "CreateMarketItem", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseCreateMarketItem(log types.Log) (*MarketplaceCreateMarketItem, error) {
	event := new(MarketplaceCreateMarketItem)
	if err := _Marketplace.contract.UnpackLog(event, "CreateMarketItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceLoseBidIterator is returned from FilterLoseBid and is used to iterate over the raw logs and unpacked data for LoseBid events raised by the Marketplace contract.
type MarketplaceLoseBidIterator struct {
	Event *MarketplaceLoseBid // Event containing the contract specifics and raw log

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
func (it *MarketplaceLoseBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceLoseBid)
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
		it.Event = new(MarketplaceLoseBid)
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
func (it *MarketplaceLoseBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceLoseBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceLoseBid represents a LoseBid event raised by the Marketplace contract.
type MarketplaceLoseBid struct {
	IdOnMarket *big.Int
	User       common.Address
	ReturnBid  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLoseBid is a free log retrieval operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_Marketplace *MarketplaceFilterer) FilterLoseBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceLoseBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceLoseBidIterator{contract: _Marketplace.contract, event: "LoseBid", logs: logs, sub: sub}, nil
}

// WatchLoseBid is a free log subscription operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_Marketplace *MarketplaceFilterer) WatchLoseBid(opts *bind.WatchOpts, sink chan<- *MarketplaceLoseBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceLoseBid)
				if err := _Marketplace.contract.UnpackLog(event, "LoseBid", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseLoseBid(log types.Log) (*MarketplaceLoseBid, error) {
	event := new(MarketplaceLoseBid)
	if err := _Marketplace.contract.UnpackLog(event, "LoseBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewAuctionPackIterator is returned from FilterNewAuctionPack and is used to iterate over the raw logs and unpacked data for NewAuctionPack events raised by the Marketplace contract.
type MarketplaceNewAuctionPackIterator struct {
	Event *MarketplaceNewAuctionPack // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewAuctionPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewAuctionPack)
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
		it.Event = new(MarketplaceNewAuctionPack)
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
func (it *MarketplaceNewAuctionPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewAuctionPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewAuctionPack represents a NewAuctionPack event raised by the Marketplace contract.
type MarketplaceNewAuctionPack struct {
	PackId            *big.Int
	StorageIdOnMarket []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewAuctionPack is a free log retrieval operation binding the contract event 0xafdb08065f8d1c171b79c618f0e85ed8641f4067ba8c4cb82583c2aa093bfa57.
//
// Solidity: event NewAuctionPack(uint256 packId, uint256[] storageIdOnMarket)
func (_Marketplace *MarketplaceFilterer) FilterNewAuctionPack(opts *bind.FilterOpts) (*MarketplaceNewAuctionPackIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewAuctionPack")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewAuctionPackIterator{contract: _Marketplace.contract, event: "NewAuctionPack", logs: logs, sub: sub}, nil
}

// WatchNewAuctionPack is a free log subscription operation binding the contract event 0xafdb08065f8d1c171b79c618f0e85ed8641f4067ba8c4cb82583c2aa093bfa57.
//
// Solidity: event NewAuctionPack(uint256 packId, uint256[] storageIdOnMarket)
func (_Marketplace *MarketplaceFilterer) WatchNewAuctionPack(opts *bind.WatchOpts, sink chan<- *MarketplaceNewAuctionPack) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewAuctionPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewAuctionPack)
				if err := _Marketplace.contract.UnpackLog(event, "NewAuctionPack", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewAuctionPack(log types.Log) (*MarketplaceNewAuctionPack, error) {
	event := new(MarketplaceNewAuctionPack)
	if err := _Marketplace.contract.UnpackLog(event, "NewAuctionPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Marketplace contract.
type MarketplaceNewBidIterator struct {
	Event *MarketplaceNewBid // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewBid)
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
		it.Event = new(MarketplaceNewBid)
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
func (it *MarketplaceNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewBid represents a NewBid event raised by the Marketplace contract.
type MarketplaceNewBid struct {
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
func (_Marketplace *MarketplaceFilterer) FilterNewBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*MarketplaceNewBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewBidIterator{contract: _Marketplace.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xc96fde502a7126831b4f59269d2f090304bc6b6158e41870ccc833175e9588f8.
//
// Solidity: event NewBid(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address bidder, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *MarketplaceNewBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewBid)
				if err := _Marketplace.contract.UnpackLog(event, "NewBid", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewBid(log types.Log) (*MarketplaceNewBid, error) {
	event := new(MarketplaceNewBid)
	if err := _Marketplace.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewFeeAddressIterator is returned from FilterNewFeeAddress and is used to iterate over the raw logs and unpacked data for NewFeeAddress events raised by the Marketplace contract.
type MarketplaceNewFeeAddressIterator struct {
	Event *MarketplaceNewFeeAddress // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewFeeAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewFeeAddress)
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
		it.Event = new(MarketplaceNewFeeAddress)
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
func (it *MarketplaceNewFeeAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewFeeAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewFeeAddress represents a NewFeeAddress event raised by the Marketplace contract.
type MarketplaceNewFeeAddress struct {
	FeeAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewFeeAddress is a free log retrieval operation binding the contract event 0xdcff558177ef1322e1a9308bd0d2b8b3922d8e05c3ec06323297bdd24db93ef0.
//
// Solidity: event NewFeeAddress(address feeAddress)
func (_Marketplace *MarketplaceFilterer) FilterNewFeeAddress(opts *bind.FilterOpts) (*MarketplaceNewFeeAddressIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewFeeAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewFeeAddressIterator{contract: _Marketplace.contract, event: "NewFeeAddress", logs: logs, sub: sub}, nil
}

// WatchNewFeeAddress is a free log subscription operation binding the contract event 0xdcff558177ef1322e1a9308bd0d2b8b3922d8e05c3ec06323297bdd24db93ef0.
//
// Solidity: event NewFeeAddress(address feeAddress)
func (_Marketplace *MarketplaceFilterer) WatchNewFeeAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceNewFeeAddress) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewFeeAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewFeeAddress)
				if err := _Marketplace.contract.UnpackLog(event, "NewFeeAddress", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewFeeAddress(log types.Log) (*MarketplaceNewFeeAddress, error) {
	event := new(MarketplaceNewFeeAddress)
	if err := _Marketplace.contract.UnpackLog(event, "NewFeeAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewFeeCreatorIterator is returned from FilterNewFeeCreator and is used to iterate over the raw logs and unpacked data for NewFeeCreator events raised by the Marketplace contract.
type MarketplaceNewFeeCreatorIterator struct {
	Event *MarketplaceNewFeeCreator // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewFeeCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewFeeCreator)
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
		it.Event = new(MarketplaceNewFeeCreator)
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
func (it *MarketplaceNewFeeCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewFeeCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewFeeCreator represents a NewFeeCreator event raised by the Marketplace contract.
type MarketplaceNewFeeCreator struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeCreator is a free log retrieval operation binding the contract event 0xefcc246bd0262a73a8223ffd295abb5ae2cab4c747b7d931061bf4123014e346.
//
// Solidity: event NewFeeCreator(uint256 percent)
func (_Marketplace *MarketplaceFilterer) FilterNewFeeCreator(opts *bind.FilterOpts) (*MarketplaceNewFeeCreatorIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewFeeCreator")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewFeeCreatorIterator{contract: _Marketplace.contract, event: "NewFeeCreator", logs: logs, sub: sub}, nil
}

// WatchNewFeeCreator is a free log subscription operation binding the contract event 0xefcc246bd0262a73a8223ffd295abb5ae2cab4c747b7d931061bf4123014e346.
//
// Solidity: event NewFeeCreator(uint256 percent)
func (_Marketplace *MarketplaceFilterer) WatchNewFeeCreator(opts *bind.WatchOpts, sink chan<- *MarketplaceNewFeeCreator) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewFeeCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewFeeCreator)
				if err := _Marketplace.contract.UnpackLog(event, "NewFeeCreator", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewFeeCreator(log types.Log) (*MarketplaceNewFeeCreator, error) {
	event := new(MarketplaceNewFeeCreator)
	if err := _Marketplace.contract.UnpackLog(event, "NewFeeCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewFeeMarketIterator is returned from FilterNewFeeMarket and is used to iterate over the raw logs and unpacked data for NewFeeMarket events raised by the Marketplace contract.
type MarketplaceNewFeeMarketIterator struct {
	Event *MarketplaceNewFeeMarket // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewFeeMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewFeeMarket)
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
		it.Event = new(MarketplaceNewFeeMarket)
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
func (it *MarketplaceNewFeeMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewFeeMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewFeeMarket represents a NewFeeMarket event raised by the Marketplace contract.
type MarketplaceNewFeeMarket struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeMarket is a free log retrieval operation binding the contract event 0xbe3b3f1e97ecc159ccad82985735211f779ad03bdd45040f03fa3fd742deab8a.
//
// Solidity: event NewFeeMarket(uint256 percent)
func (_Marketplace *MarketplaceFilterer) FilterNewFeeMarket(opts *bind.FilterOpts) (*MarketplaceNewFeeMarketIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewFeeMarket")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewFeeMarketIterator{contract: _Marketplace.contract, event: "NewFeeMarket", logs: logs, sub: sub}, nil
}

// WatchNewFeeMarket is a free log subscription operation binding the contract event 0xbe3b3f1e97ecc159ccad82985735211f779ad03bdd45040f03fa3fd742deab8a.
//
// Solidity: event NewFeeMarket(uint256 percent)
func (_Marketplace *MarketplaceFilterer) WatchNewFeeMarket(opts *bind.WatchOpts, sink chan<- *MarketplaceNewFeeMarket) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewFeeMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewFeeMarket)
				if err := _Marketplace.contract.UnpackLog(event, "NewFeeMarket", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewFeeMarket(log types.Log) (*MarketplaceNewFeeMarket, error) {
	event := new(MarketplaceNewFeeMarket)
	if err := _Marketplace.contract.UnpackLog(event, "NewFeeMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewFeeRefIterator is returned from FilterNewFeeRef and is used to iterate over the raw logs and unpacked data for NewFeeRef events raised by the Marketplace contract.
type MarketplaceNewFeeRefIterator struct {
	Event *MarketplaceNewFeeRef // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewFeeRefIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewFeeRef)
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
		it.Event = new(MarketplaceNewFeeRef)
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
func (it *MarketplaceNewFeeRefIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewFeeRefIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewFeeRef represents a NewFeeRef event raised by the Marketplace contract.
type MarketplaceNewFeeRef struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRef is a free log retrieval operation binding the contract event 0x09c0a82bf115f44d83933f5baafce9d671ef57977cdd2f45487e22fbe5c9b47d.
//
// Solidity: event NewFeeRef(uint256 percent)
func (_Marketplace *MarketplaceFilterer) FilterNewFeeRef(opts *bind.FilterOpts) (*MarketplaceNewFeeRefIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewFeeRef")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewFeeRefIterator{contract: _Marketplace.contract, event: "NewFeeRef", logs: logs, sub: sub}, nil
}

// WatchNewFeeRef is a free log subscription operation binding the contract event 0x09c0a82bf115f44d83933f5baafce9d671ef57977cdd2f45487e22fbe5c9b47d.
//
// Solidity: event NewFeeRef(uint256 percent)
func (_Marketplace *MarketplaceFilterer) WatchNewFeeRef(opts *bind.WatchOpts, sink chan<- *MarketplaceNewFeeRef) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewFeeRef")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewFeeRef)
				if err := _Marketplace.contract.UnpackLog(event, "NewFeeRef", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewFeeRef(log types.Log) (*MarketplaceNewFeeRef, error) {
	event := new(MarketplaceNewFeeRef)
	if err := _Marketplace.contract.UnpackLog(event, "NewFeeRef", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewMainTokenAddressIterator is returned from FilterNewMainTokenAddress and is used to iterate over the raw logs and unpacked data for NewMainTokenAddress events raised by the Marketplace contract.
type MarketplaceNewMainTokenAddressIterator struct {
	Event *MarketplaceNewMainTokenAddress // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewMainTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewMainTokenAddress)
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
		it.Event = new(MarketplaceNewMainTokenAddress)
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
func (it *MarketplaceNewMainTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewMainTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewMainTokenAddress represents a NewMainTokenAddress event raised by the Marketplace contract.
type MarketplaceNewMainTokenAddress struct {
	MainToken common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewMainTokenAddress is a free log retrieval operation binding the contract event 0xb560ea6db584a0e0890a534beed59ff9f580da1f25be95f51dd1d580e335e021.
//
// Solidity: event NewMainTokenAddress(address mainToken)
func (_Marketplace *MarketplaceFilterer) FilterNewMainTokenAddress(opts *bind.FilterOpts) (*MarketplaceNewMainTokenAddressIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewMainTokenAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewMainTokenAddressIterator{contract: _Marketplace.contract, event: "NewMainTokenAddress", logs: logs, sub: sub}, nil
}

// WatchNewMainTokenAddress is a free log subscription operation binding the contract event 0xb560ea6db584a0e0890a534beed59ff9f580da1f25be95f51dd1d580e335e021.
//
// Solidity: event NewMainTokenAddress(address mainToken)
func (_Marketplace *MarketplaceFilterer) WatchNewMainTokenAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceNewMainTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewMainTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewMainTokenAddress)
				if err := _Marketplace.contract.UnpackLog(event, "NewMainTokenAddress", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewMainTokenAddress(log types.Log) (*MarketplaceNewMainTokenAddress, error) {
	event := new(MarketplaceNewMainTokenAddress)
	if err := _Marketplace.contract.UnpackLog(event, "NewMainTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewPackIterator is returned from FilterNewPack and is used to iterate over the raw logs and unpacked data for NewPack events raised by the Marketplace contract.
type MarketplaceNewPackIterator struct {
	Event *MarketplaceNewPack // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewPack)
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
		it.Event = new(MarketplaceNewPack)
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
func (it *MarketplaceNewPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewPack represents a NewPack event raised by the Marketplace contract.
type MarketplaceNewPack struct {
	Pack        *big.Int
	IdsOnMarket []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPack is a free log retrieval operation binding the contract event 0x46edc46b662430a962a47432ce24ef4cbf0e481e2b27a5f93fec3ea0c52d8887.
//
// Solidity: event NewPack(uint256 pack, uint256[] idsOnMarket)
func (_Marketplace *MarketplaceFilterer) FilterNewPack(opts *bind.FilterOpts) (*MarketplaceNewPackIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewPack")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewPackIterator{contract: _Marketplace.contract, event: "NewPack", logs: logs, sub: sub}, nil
}

// WatchNewPack is a free log subscription operation binding the contract event 0x46edc46b662430a962a47432ce24ef4cbf0e481e2b27a5f93fec3ea0c52d8887.
//
// Solidity: event NewPack(uint256 pack, uint256[] idsOnMarket)
func (_Marketplace *MarketplaceFilterer) WatchNewPack(opts *bind.WatchOpts, sink chan<- *MarketplaceNewPack) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewPack)
				if err := _Marketplace.contract.UnpackLog(event, "NewPack", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewPack(log types.Log) (*MarketplaceNewPack, error) {
	event := new(MarketplaceNewPack)
	if err := _Marketplace.contract.UnpackLog(event, "NewPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewRefAddressIterator is returned from FilterNewRefAddress and is used to iterate over the raw logs and unpacked data for NewRefAddress events raised by the Marketplace contract.
type MarketplaceNewRefAddressIterator struct {
	Event *MarketplaceNewRefAddress // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewRefAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewRefAddress)
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
		it.Event = new(MarketplaceNewRefAddress)
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
func (it *MarketplaceNewRefAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewRefAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewRefAddress represents a NewRefAddress event raised by the Marketplace contract.
type MarketplaceNewRefAddress struct {
	RefAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRefAddress is a free log retrieval operation binding the contract event 0x7e759132c178bdc3b8f23494e9c61700d192b9fcd42420e58a82811c1e4f8d54.
//
// Solidity: event NewRefAddress(address refAddress)
func (_Marketplace *MarketplaceFilterer) FilterNewRefAddress(opts *bind.FilterOpts) (*MarketplaceNewRefAddressIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewRefAddress")
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewRefAddressIterator{contract: _Marketplace.contract, event: "NewRefAddress", logs: logs, sub: sub}, nil
}

// WatchNewRefAddress is a free log subscription operation binding the contract event 0x7e759132c178bdc3b8f23494e9c61700d192b9fcd42420e58a82811c1e4f8d54.
//
// Solidity: event NewRefAddress(address refAddress)
func (_Marketplace *MarketplaceFilterer) WatchNewRefAddress(opts *bind.WatchOpts, sink chan<- *MarketplaceNewRefAddress) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewRefAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewRefAddress)
				if err := _Marketplace.contract.UnpackLog(event, "NewRefAddress", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseNewRefAddress(log types.Log) (*MarketplaceNewRefAddress, error) {
	event := new(MarketplaceNewRefAddress)
	if err := _Marketplace.contract.UnpackLog(event, "NewRefAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRemoveBlackListFeeIterator is returned from FilterRemoveBlackListFee and is used to iterate over the raw logs and unpacked data for RemoveBlackListFee events raised by the Marketplace contract.
type MarketplaceRemoveBlackListFeeIterator struct {
	Event *MarketplaceRemoveBlackListFee // Event containing the contract specifics and raw log

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
func (it *MarketplaceRemoveBlackListFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRemoveBlackListFee)
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
		it.Event = new(MarketplaceRemoveBlackListFee)
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
func (it *MarketplaceRemoveBlackListFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRemoveBlackListFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRemoveBlackListFee represents a RemoveBlackListFee event raised by the Marketplace contract.
type MarketplaceRemoveBlackListFee struct {
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlackListFee is a free log retrieval operation binding the contract event 0x2c3069ddf8f5af7e3a769d0f83893e40875365e3902dab291c5a910b44dced14.
//
// Solidity: event RemoveBlackListFee(address[] users)
func (_Marketplace *MarketplaceFilterer) FilterRemoveBlackListFee(opts *bind.FilterOpts) (*MarketplaceRemoveBlackListFeeIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RemoveBlackListFee")
	if err != nil {
		return nil, err
	}
	return &MarketplaceRemoveBlackListFeeIterator{contract: _Marketplace.contract, event: "RemoveBlackListFee", logs: logs, sub: sub}, nil
}

// WatchRemoveBlackListFee is a free log subscription operation binding the contract event 0x2c3069ddf8f5af7e3a769d0f83893e40875365e3902dab291c5a910b44dced14.
//
// Solidity: event RemoveBlackListFee(address[] users)
func (_Marketplace *MarketplaceFilterer) WatchRemoveBlackListFee(opts *bind.WatchOpts, sink chan<- *MarketplaceRemoveBlackListFee) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RemoveBlackListFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRemoveBlackListFee)
				if err := _Marketplace.contract.UnpackLog(event, "RemoveBlackListFee", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRemoveBlackListFee(log types.Log) (*MarketplaceRemoveBlackListFee, error) {
	event := new(MarketplaceRemoveBlackListFee)
	if err := _Marketplace.contract.UnpackLog(event, "RemoveBlackListFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRemoveFactoryBasicIterator is returned from FilterRemoveFactoryBasic and is used to iterate over the raw logs and unpacked data for RemoveFactoryBasic events raised by the Marketplace contract.
type MarketplaceRemoveFactoryBasicIterator struct {
	Event *MarketplaceRemoveFactoryBasic // Event containing the contract specifics and raw log

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
func (it *MarketplaceRemoveFactoryBasicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRemoveFactoryBasic)
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
		it.Event = new(MarketplaceRemoveFactoryBasic)
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
func (it *MarketplaceRemoveFactoryBasicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRemoveFactoryBasicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRemoveFactoryBasic represents a RemoveFactoryBasic event raised by the Marketplace contract.
type MarketplaceRemoveFactoryBasic struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveFactoryBasic is a free log retrieval operation binding the contract event 0x4f88ae8a1a0a602ddff03d4a85df29951abd94bb171d32b6825fc55bdeecc7bf.
//
// Solidity: event RemoveFactoryBasic(address factory)
func (_Marketplace *MarketplaceFilterer) FilterRemoveFactoryBasic(opts *bind.FilterOpts) (*MarketplaceRemoveFactoryBasicIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RemoveFactoryBasic")
	if err != nil {
		return nil, err
	}
	return &MarketplaceRemoveFactoryBasicIterator{contract: _Marketplace.contract, event: "RemoveFactoryBasic", logs: logs, sub: sub}, nil
}

// WatchRemoveFactoryBasic is a free log subscription operation binding the contract event 0x4f88ae8a1a0a602ddff03d4a85df29951abd94bb171d32b6825fc55bdeecc7bf.
//
// Solidity: event RemoveFactoryBasic(address factory)
func (_Marketplace *MarketplaceFilterer) WatchRemoveFactoryBasic(opts *bind.WatchOpts, sink chan<- *MarketplaceRemoveFactoryBasic) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RemoveFactoryBasic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRemoveFactoryBasic)
				if err := _Marketplace.contract.UnpackLog(event, "RemoveFactoryBasic", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRemoveFactoryBasic(log types.Log) (*MarketplaceRemoveFactoryBasic, error) {
	event := new(MarketplaceRemoveFactoryBasic)
	if err := _Marketplace.contract.UnpackLog(event, "RemoveFactoryBasic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRemoveFactoryVipIterator is returned from FilterRemoveFactoryVip and is used to iterate over the raw logs and unpacked data for RemoveFactoryVip events raised by the Marketplace contract.
type MarketplaceRemoveFactoryVipIterator struct {
	Event *MarketplaceRemoveFactoryVip // Event containing the contract specifics and raw log

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
func (it *MarketplaceRemoveFactoryVipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRemoveFactoryVip)
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
		it.Event = new(MarketplaceRemoveFactoryVip)
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
func (it *MarketplaceRemoveFactoryVipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRemoveFactoryVipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRemoveFactoryVip represents a RemoveFactoryVip event raised by the Marketplace contract.
type MarketplaceRemoveFactoryVip struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveFactoryVip is a free log retrieval operation binding the contract event 0xef7affc8c54e6b9a0413bdc46ec5c49dd3ecd035a5e6763f8c0fac4cd256ff64.
//
// Solidity: event RemoveFactoryVip(address factory)
func (_Marketplace *MarketplaceFilterer) FilterRemoveFactoryVip(opts *bind.FilterOpts) (*MarketplaceRemoveFactoryVipIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RemoveFactoryVip")
	if err != nil {
		return nil, err
	}
	return &MarketplaceRemoveFactoryVipIterator{contract: _Marketplace.contract, event: "RemoveFactoryVip", logs: logs, sub: sub}, nil
}

// WatchRemoveFactoryVip is a free log subscription operation binding the contract event 0xef7affc8c54e6b9a0413bdc46ec5c49dd3ecd035a5e6763f8c0fac4cd256ff64.
//
// Solidity: event RemoveFactoryVip(address factory)
func (_Marketplace *MarketplaceFilterer) WatchRemoveFactoryVip(opts *bind.WatchOpts, sink chan<- *MarketplaceRemoveFactoryVip) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RemoveFactoryVip")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRemoveFactoryVip)
				if err := _Marketplace.contract.UnpackLog(event, "RemoveFactoryVip", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRemoveFactoryVip(log types.Log) (*MarketplaceRemoveFactoryVip, error) {
	event := new(MarketplaceRemoveFactoryVip)
	if err := _Marketplace.contract.UnpackLog(event, "RemoveFactoryVip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
