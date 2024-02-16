// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

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

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CancelSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ClaimNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"}],\"name\":\"CreateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnBid\",\"type\":\"uint256\"}],\"name\":\"LoseBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idOnMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AllFactoryBasic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AllFactoryVip\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Factory\",\"outputs\":[{\"internalType\":\"contractERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MainToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RefContract\",\"outputs\":[{\"internalType\":\"contractOPV_REF\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blackListFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"cancelSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idOnMarket\",\"type\":\"uint256\"}],\"name\":\"claimNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"getBuyData\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeCreator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRef\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRefAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MainToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_RefAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_FeeAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"removeBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryBasic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeFactoryVip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractHold\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"saleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"}],\"name\":\"setBlackListFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setFeeRef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MainToken\",\"type\":\"address\"}],\"name\":\"setMainTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_RefAddress\",\"type\":\"address\"}],\"name\":\"setRefAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// AllFactoryBasic is a free data retrieval call binding the contract method 0x6b80b4ac.
//
// Solidity: function AllFactoryBasic(address ) view returns(bool)
func (_Auction *AuctionCaller) AllFactoryBasic(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "AllFactoryBasic", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllFactoryBasic is a free data retrieval call binding the contract method 0x6b80b4ac.
//
// Solidity: function AllFactoryBasic(address ) view returns(bool)
func (_Auction *AuctionSession) AllFactoryBasic(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllFactoryBasic(&_Auction.CallOpts, arg0)
}

// AllFactoryBasic is a free data retrieval call binding the contract method 0x6b80b4ac.
//
// Solidity: function AllFactoryBasic(address ) view returns(bool)
func (_Auction *AuctionCallerSession) AllFactoryBasic(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllFactoryBasic(&_Auction.CallOpts, arg0)
}

// AllFactoryVip is a free data retrieval call binding the contract method 0xe340442b.
//
// Solidity: function AllFactoryVip(address ) view returns(bool)
func (_Auction *AuctionCaller) AllFactoryVip(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "AllFactoryVip", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllFactoryVip is a free data retrieval call binding the contract method 0xe340442b.
//
// Solidity: function AllFactoryVip(address ) view returns(bool)
func (_Auction *AuctionSession) AllFactoryVip(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllFactoryVip(&_Auction.CallOpts, arg0)
}

// AllFactoryVip is a free data retrieval call binding the contract method 0xe340442b.
//
// Solidity: function AllFactoryVip(address ) view returns(bool)
func (_Auction *AuctionCallerSession) AllFactoryVip(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllFactoryVip(&_Auction.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_Auction *AuctionCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_Auction *AuctionSession) Factory() (common.Address, error) {
	return _Auction.Contract.Factory(&_Auction.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_Auction *AuctionCallerSession) Factory() (common.Address, error) {
	return _Auction.Contract.Factory(&_Auction.CallOpts)
}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_Auction *AuctionCaller) MainToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "MainToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_Auction *AuctionSession) MainToken() (common.Address, error) {
	return _Auction.Contract.MainToken(&_Auction.CallOpts)
}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_Auction *AuctionCallerSession) MainToken() (common.Address, error) {
	return _Auction.Contract.MainToken(&_Auction.CallOpts)
}

// RefContract is a free data retrieval call binding the contract method 0x6781c09e.
//
// Solidity: function RefContract() view returns(address)
func (_Auction *AuctionCaller) RefContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "RefContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RefContract is a free data retrieval call binding the contract method 0x6781c09e.
//
// Solidity: function RefContract() view returns(address)
func (_Auction *AuctionSession) RefContract() (common.Address, error) {
	return _Auction.Contract.RefContract(&_Auction.CallOpts)
}

// RefContract is a free data retrieval call binding the contract method 0x6781c09e.
//
// Solidity: function RefContract() view returns(address)
func (_Auction *AuctionCallerSession) RefContract() (common.Address, error) {
	return _Auction.Contract.RefContract(&_Auction.CallOpts)
}

// BlackListFee is a free data retrieval call binding the contract method 0xa415499b.
//
// Solidity: function blackListFee(address ) view returns(bool)
func (_Auction *AuctionCaller) BlackListFee(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "blackListFee", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackListFee is a free data retrieval call binding the contract method 0xa415499b.
//
// Solidity: function blackListFee(address ) view returns(bool)
func (_Auction *AuctionSession) BlackListFee(arg0 common.Address) (bool, error) {
	return _Auction.Contract.BlackListFee(&_Auction.CallOpts, arg0)
}

// BlackListFee is a free data retrieval call binding the contract method 0xa415499b.
//
// Solidity: function blackListFee(address ) view returns(bool)
func (_Auction *AuctionCallerSession) BlackListFee(arg0 common.Address) (bool, error) {
	return _Auction.Contract.BlackListFee(&_Auction.CallOpts, arg0)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_Auction *AuctionCaller) GetBuyData(opts *bind.CallOpts, price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getBuyData", price_, factory_, seller_, tokenId_)

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
func (_Auction *AuctionSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _Auction.Contract.GetBuyData(&_Auction.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetBuyData is a free data retrieval call binding the contract method 0xe4984d32.
//
// Solidity: function getBuyData(uint256 price_, address factory_, address seller_, uint256 tokenId_) view returns(uint256[], address[])
func (_Auction *AuctionCallerSession) GetBuyData(price_ *big.Int, factory_ common.Address, seller_ common.Address, tokenId_ *big.Int) ([]*big.Int, []common.Address, error) {
	return _Auction.Contract.GetBuyData(&_Auction.CallOpts, price_, factory_, seller_, tokenId_)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Auction *AuctionCaller) GetFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Auction *AuctionSession) GetFeeAddress() (common.Address, error) {
	return _Auction.Contract.GetFeeAddress(&_Auction.CallOpts)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_Auction *AuctionCallerSession) GetFeeAddress() (common.Address, error) {
	return _Auction.Contract.GetFeeAddress(&_Auction.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Auction *AuctionCaller) GetFeeCreator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getFeeCreator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Auction *AuctionSession) GetFeeCreator() (*big.Int, error) {
	return _Auction.Contract.GetFeeCreator(&_Auction.CallOpts)
}

// GetFeeCreator is a free data retrieval call binding the contract method 0x4566a9ce.
//
// Solidity: function getFeeCreator() view returns(uint256)
func (_Auction *AuctionCallerSession) GetFeeCreator() (*big.Int, error) {
	return _Auction.Contract.GetFeeCreator(&_Auction.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Auction *AuctionCaller) GetFeeMarket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getFeeMarket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Auction *AuctionSession) GetFeeMarket() (*big.Int, error) {
	return _Auction.Contract.GetFeeMarket(&_Auction.CallOpts)
}

// GetFeeMarket is a free data retrieval call binding the contract method 0x2e8eedd1.
//
// Solidity: function getFeeMarket() view returns(uint256)
func (_Auction *AuctionCallerSession) GetFeeMarket() (*big.Int, error) {
	return _Auction.Contract.GetFeeMarket(&_Auction.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Auction *AuctionCaller) GetFeeRef(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getFeeRef")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Auction *AuctionSession) GetFeeRef() (*big.Int, error) {
	return _Auction.Contract.GetFeeRef(&_Auction.CallOpts)
}

// GetFeeRef is a free data retrieval call binding the contract method 0x794d1d4c.
//
// Solidity: function getFeeRef() view returns(uint256)
func (_Auction *AuctionCallerSession) GetFeeRef() (*big.Int, error) {
	return _Auction.Contract.GetFeeRef(&_Auction.CallOpts)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Auction *AuctionCaller) GetRefAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getRefAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Auction *AuctionSession) GetRefAddress() (common.Address, error) {
	return _Auction.Contract.GetRefAddress(&_Auction.CallOpts)
}

// GetRefAddress is a free data retrieval call binding the contract method 0x3f9ce517.
//
// Solidity: function getRefAddress() view returns(address)
func (_Auction *AuctionCallerSession) GetRefAddress() (common.Address, error) {
	return _Auction.Contract.GetRefAddress(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts, _idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "Bid", _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Auction *AuctionSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, _idOnMarket, _price)
}

// Bid is a paid mutator transaction binding the contract method 0x47a0375b.
//
// Solidity: function Bid(uint256 _idOnMarket, uint256 _price) returns()
func (_Auction *AuctionTransactorSession) Bid(_idOnMarket *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, _idOnMarket, _price)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Auction *AuctionTransactor) AddFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "addFactoryBasic", proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Auction *AuctionSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.AddFactoryBasic(&_Auction.TransactOpts, proxy)
}

// AddFactoryBasic is a paid mutator transaction binding the contract method 0xada74f02.
//
// Solidity: function addFactoryBasic(address proxy) returns()
func (_Auction *AuctionTransactorSession) AddFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.AddFactoryBasic(&_Auction.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Auction *AuctionTransactor) AddFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "addFactoryVip", proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Auction *AuctionSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.AddFactoryVip(&_Auction.TransactOpts, proxy)
}

// AddFactoryVip is a paid mutator transaction binding the contract method 0x445bc28e.
//
// Solidity: function addFactoryVip(address proxy) returns()
func (_Auction *AuctionTransactorSession) AddFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.AddFactoryVip(&_Auction.TransactOpts, proxy)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Auction *AuctionTransactor) CancelSell(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "cancelSell", _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Auction *AuctionSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.CancelSell(&_Auction.TransactOpts, _idOnMarket)
}

// CancelSell is a paid mutator transaction binding the contract method 0x07c9cd45.
//
// Solidity: function cancelSell(uint256 _idOnMarket) returns()
func (_Auction *AuctionTransactorSession) CancelSell(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.CancelSell(&_Auction.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Auction *AuctionTransactor) ClaimNft(opts *bind.TransactOpts, _idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "claimNft", _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Auction *AuctionSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimNft(&_Auction.TransactOpts, _idOnMarket)
}

// ClaimNft is a paid mutator transaction binding the contract method 0x2ec09d39.
//
// Solidity: function claimNft(uint256 _idOnMarket) returns()
func (_Auction *AuctionTransactorSession) ClaimNft(_idOnMarket *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimNft(&_Auction.TransactOpts, _idOnMarket)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _MainToken, address _RefAddress, address _FeeAddress) returns()
func (_Auction *AuctionTransactor) Initialize(opts *bind.TransactOpts, _MainToken common.Address, _RefAddress common.Address, _FeeAddress common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "initialize", _MainToken, _RefAddress, _FeeAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _MainToken, address _RefAddress, address _FeeAddress) returns()
func (_Auction *AuctionSession) Initialize(_MainToken common.Address, _RefAddress common.Address, _FeeAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _MainToken, _RefAddress, _FeeAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _MainToken, address _RefAddress, address _FeeAddress) returns()
func (_Auction *AuctionTransactorSession) Initialize(_MainToken common.Address, _RefAddress common.Address, _FeeAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _MainToken, _RefAddress, _FeeAddress)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Auction *AuctionTransactor) RemoveBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "removeBlackListFee", user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Auction *AuctionSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveBlackListFee(&_Auction.TransactOpts, user)
}

// RemoveBlackListFee is a paid mutator transaction binding the contract method 0x972bac07.
//
// Solidity: function removeBlackListFee(address[] user) returns()
func (_Auction *AuctionTransactorSession) RemoveBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveBlackListFee(&_Auction.TransactOpts, user)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Auction *AuctionTransactor) RemoveFactoryBasic(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "removeFactoryBasic", proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Auction *AuctionSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveFactoryBasic(&_Auction.TransactOpts, proxy)
}

// RemoveFactoryBasic is a paid mutator transaction binding the contract method 0x1b867c9c.
//
// Solidity: function removeFactoryBasic(address proxy) returns()
func (_Auction *AuctionTransactorSession) RemoveFactoryBasic(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveFactoryBasic(&_Auction.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Auction *AuctionTransactor) RemoveFactoryVip(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "removeFactoryVip", proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Auction *AuctionSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveFactoryVip(&_Auction.TransactOpts, proxy)
}

// RemoveFactoryVip is a paid mutator transaction binding the contract method 0x6567bfcd.
//
// Solidity: function removeFactoryVip(address proxy) returns()
func (_Auction *AuctionTransactorSession) RemoveFactoryVip(proxy common.Address) (*types.Transaction, error) {
	return _Auction.Contract.RemoveFactoryVip(&_Auction.TransactOpts, proxy)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x0625c937.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, address _contractHold, address _factory) returns()
func (_Auction *AuctionTransactor) SaleAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "saleAuction", _tokenId, _startTime, _endTime, _minBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x0625c937.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, address _contractHold, address _factory) returns()
func (_Auction *AuctionSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SaleAuction(&_Auction.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _contractHold, _factory)
}

// SaleAuction is a paid mutator transaction binding the contract method 0x0625c937.
//
// Solidity: function saleAuction(uint256 _tokenId, uint256 _startTime, uint256 _endTime, uint256 _minBid, address _contractHold, address _factory) returns()
func (_Auction *AuctionTransactorSession) SaleAuction(_tokenId *big.Int, _startTime *big.Int, _endTime *big.Int, _minBid *big.Int, _contractHold common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SaleAuction(&_Auction.TransactOpts, _tokenId, _startTime, _endTime, _minBid, _contractHold, _factory)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Auction *AuctionTransactor) SetBlackListFee(opts *bind.TransactOpts, user []common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setBlackListFee", user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Auction *AuctionSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetBlackListFee(&_Auction.TransactOpts, user)
}

// SetBlackListFee is a paid mutator transaction binding the contract method 0x235068fc.
//
// Solidity: function setBlackListFee(address[] user) returns()
func (_Auction *AuctionTransactorSession) SetBlackListFee(user []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetBlackListFee(&_Auction.TransactOpts, user)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Auction *AuctionTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Auction *AuctionSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeAddress(&_Auction.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Auction *AuctionTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeAddress(&_Auction.TransactOpts, _feeAddress)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Auction *AuctionTransactor) SetFeeCreator(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setFeeCreator", percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Auction *AuctionSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeCreator(&_Auction.TransactOpts, percent)
}

// SetFeeCreator is a paid mutator transaction binding the contract method 0x50885acc.
//
// Solidity: function setFeeCreator(uint256 percent) returns()
func (_Auction *AuctionTransactorSession) SetFeeCreator(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeCreator(&_Auction.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Auction *AuctionTransactor) SetFeeMarket(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setFeeMarket", percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Auction *AuctionSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeMarket(&_Auction.TransactOpts, percent)
}

// SetFeeMarket is a paid mutator transaction binding the contract method 0x48eddc38.
//
// Solidity: function setFeeMarket(uint256 percent) returns()
func (_Auction *AuctionTransactorSession) SetFeeMarket(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeMarket(&_Auction.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Auction *AuctionTransactor) SetFeeRef(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setFeeRef", percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Auction *AuctionSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeRef(&_Auction.TransactOpts, percent)
}

// SetFeeRef is a paid mutator transaction binding the contract method 0xaacde203.
//
// Solidity: function setFeeRef(uint256 percent) returns()
func (_Auction *AuctionTransactorSession) SetFeeRef(percent *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetFeeRef(&_Auction.TransactOpts, percent)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Auction *AuctionTransactor) SetMainTokenAddress(opts *bind.TransactOpts, _MainToken common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setMainTokenAddress", _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Auction *AuctionSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetMainTokenAddress(&_Auction.TransactOpts, _MainToken)
}

// SetMainTokenAddress is a paid mutator transaction binding the contract method 0x6809fedd.
//
// Solidity: function setMainTokenAddress(address _MainToken) returns()
func (_Auction *AuctionTransactorSession) SetMainTokenAddress(_MainToken common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetMainTokenAddress(&_Auction.TransactOpts, _MainToken)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Auction *AuctionTransactor) SetRefAddress(opts *bind.TransactOpts, _RefAddress common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setRefAddress", _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Auction *AuctionSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetRefAddress(&_Auction.TransactOpts, _RefAddress)
}

// SetRefAddress is a paid mutator transaction binding the contract method 0x061e26af.
//
// Solidity: function setRefAddress(address _RefAddress) returns()
func (_Auction *AuctionTransactorSession) SetRefAddress(_RefAddress common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetRefAddress(&_Auction.TransactOpts, _RefAddress)
}

// AuctionCancelSellIterator is returned from FilterCancelSell and is used to iterate over the raw logs and unpacked data for CancelSell events raised by the Auction contract.
type AuctionCancelSellIterator struct {
	Event *AuctionCancelSell // Event containing the contract specifics and raw log

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
func (it *AuctionCancelSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionCancelSell)
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
		it.Event = new(AuctionCancelSell)
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
func (it *AuctionCancelSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionCancelSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionCancelSell represents a CancelSell event raised by the Auction contract.
type AuctionCancelSell struct {
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
func (_Auction *AuctionFilterer) FilterCancelSell(opts *bind.FilterOpts, idOnMarket []*big.Int) (*AuctionCancelSellIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &AuctionCancelSellIterator{contract: _Auction.contract, event: "CancelSell", logs: logs, sub: sub}, nil
}

// WatchCancelSell is a free log subscription operation binding the contract event 0x29a6902ae24fe755451b145d47596d0840bfbe1c5044ecee155bf4aec677fa02.
//
// Solidity: event CancelSell(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address owner)
func (_Auction *AuctionFilterer) WatchCancelSell(opts *bind.WatchOpts, sink chan<- *AuctionCancelSell, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "CancelSell", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionCancelSell)
				if err := _Auction.contract.UnpackLog(event, "CancelSell", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseCancelSell(log types.Log) (*AuctionCancelSell, error) {
	event := new(AuctionCancelSell)
	if err := _Auction.contract.UnpackLog(event, "CancelSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionClaimNftIterator is returned from FilterClaimNft and is used to iterate over the raw logs and unpacked data for ClaimNft events raised by the Auction contract.
type AuctionClaimNftIterator struct {
	Event *AuctionClaimNft // Event containing the contract specifics and raw log

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
func (it *AuctionClaimNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionClaimNft)
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
		it.Event = new(AuctionClaimNft)
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
func (it *AuctionClaimNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionClaimNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionClaimNft represents a ClaimNft event raised by the Auction contract.
type AuctionClaimNft struct {
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
func (_Auction *AuctionFilterer) FilterClaimNft(opts *bind.FilterOpts, idOnMarket []*big.Int) (*AuctionClaimNftIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &AuctionClaimNftIterator{contract: _Auction.contract, event: "ClaimNft", logs: logs, sub: sub}, nil
}

// WatchClaimNft is a free log subscription operation binding the contract event 0xad134c2ad1586b6c1b233ec09b85165f2e35b759c78a884e5f2c8dda9211f2b4.
//
// Solidity: event ClaimNft(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address buyer, uint256 price)
func (_Auction *AuctionFilterer) WatchClaimNft(opts *bind.WatchOpts, sink chan<- *AuctionClaimNft, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "ClaimNft", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionClaimNft)
				if err := _Auction.contract.UnpackLog(event, "ClaimNft", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseClaimNft(log types.Log) (*AuctionClaimNft, error) {
	event := new(AuctionClaimNft)
	if err := _Auction.contract.UnpackLog(event, "ClaimNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionCreateAuctionIterator is returned from FilterCreateAuction and is used to iterate over the raw logs and unpacked data for CreateAuction events raised by the Auction contract.
type AuctionCreateAuctionIterator struct {
	Event *AuctionCreateAuction // Event containing the contract specifics and raw log

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
func (it *AuctionCreateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionCreateAuction)
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
		it.Event = new(AuctionCreateAuction)
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
func (it *AuctionCreateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionCreateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionCreateAuction represents a CreateAuction event raised by the Auction contract.
type AuctionCreateAuction struct {
	IdOnMarket  *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	MinBid      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateAuction is a free log retrieval operation binding the contract event 0xf40f0d6dd3179c8516363485d2a69e02d5f5ff8ecd743cb0f8ecb3cbb6cdaf17.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid)
func (_Auction *AuctionFilterer) FilterCreateAuction(opts *bind.FilterOpts, idOnMarket []*big.Int) (*AuctionCreateAuctionIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &AuctionCreateAuctionIterator{contract: _Auction.contract, event: "CreateAuction", logs: logs, sub: sub}, nil
}

// WatchCreateAuction is a free log subscription operation binding the contract event 0xf40f0d6dd3179c8516363485d2a69e02d5f5ff8ecd743cb0f8ecb3cbb6cdaf17.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid)
func (_Auction *AuctionFilterer) WatchCreateAuction(opts *bind.WatchOpts, sink chan<- *AuctionCreateAuction, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "CreateAuction", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionCreateAuction)
				if err := _Auction.contract.UnpackLog(event, "CreateAuction", log); err != nil {
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

// ParseCreateAuction is a log parse operation binding the contract event 0xf40f0d6dd3179c8516363485d2a69e02d5f5ff8ecd743cb0f8ecb3cbb6cdaf17.
//
// Solidity: event CreateAuction(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, uint256 _startTime, uint256 _endTime, uint256 _minBid)
func (_Auction *AuctionFilterer) ParseCreateAuction(log types.Log) (*AuctionCreateAuction, error) {
	event := new(AuctionCreateAuction)
	if err := _Auction.contract.UnpackLog(event, "CreateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionLoseBidIterator is returned from FilterLoseBid and is used to iterate over the raw logs and unpacked data for LoseBid events raised by the Auction contract.
type AuctionLoseBidIterator struct {
	Event *AuctionLoseBid // Event containing the contract specifics and raw log

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
func (it *AuctionLoseBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionLoseBid)
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
		it.Event = new(AuctionLoseBid)
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
func (it *AuctionLoseBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionLoseBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionLoseBid represents a LoseBid event raised by the Auction contract.
type AuctionLoseBid struct {
	IdOnMarket *big.Int
	User       common.Address
	ReturnBid  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLoseBid is a free log retrieval operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_Auction *AuctionFilterer) FilterLoseBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*AuctionLoseBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &AuctionLoseBidIterator{contract: _Auction.contract, event: "LoseBid", logs: logs, sub: sub}, nil
}

// WatchLoseBid is a free log subscription operation binding the contract event 0xbf6d98994b85929aeaa0a8c7eb0b4e33205def8d0dcb222ff2da236e4efb3f05.
//
// Solidity: event LoseBid(uint256 indexed idOnMarket, address user, uint256 returnBid)
func (_Auction *AuctionFilterer) WatchLoseBid(opts *bind.WatchOpts, sink chan<- *AuctionLoseBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "LoseBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionLoseBid)
				if err := _Auction.contract.UnpackLog(event, "LoseBid", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseLoseBid(log types.Log) (*AuctionLoseBid, error) {
	event := new(AuctionLoseBid)
	if err := _Auction.contract.UnpackLog(event, "LoseBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Auction contract.
type AuctionNewBidIterator struct {
	Event *AuctionNewBid // Event containing the contract specifics and raw log

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
func (it *AuctionNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionNewBid)
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
		it.Event = new(AuctionNewBid)
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
func (it *AuctionNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionNewBid represents a NewBid event raised by the Auction contract.
type AuctionNewBid struct {
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
func (_Auction *AuctionFilterer) FilterNewBid(opts *bind.FilterOpts, idOnMarket []*big.Int) (*AuctionNewBidIterator, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return &AuctionNewBidIterator{contract: _Auction.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xc96fde502a7126831b4f59269d2f090304bc6b6158e41870ccc833175e9588f8.
//
// Solidity: event NewBid(uint256 indexed idOnMarket, uint256 tokenId, address nftContract, address seller, address bidder, uint256 price)
func (_Auction *AuctionFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *AuctionNewBid, idOnMarket []*big.Int) (event.Subscription, error) {

	var idOnMarketRule []interface{}
	for _, idOnMarketItem := range idOnMarket {
		idOnMarketRule = append(idOnMarketRule, idOnMarketItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "NewBid", idOnMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionNewBid)
				if err := _Auction.contract.UnpackLog(event, "NewBid", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseNewBid(log types.Log) (*AuctionNewBid, error) {
	event := new(AuctionNewBid)
	if err := _Auction.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
