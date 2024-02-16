// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package create_event_sec

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

// CreateEventMetaData contains all meta data concerning the CreateEvent contract.
var CreateEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundToSystem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundToRef\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFund\",\"type\":\"uint256\"}],\"name\":\"JoinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"OutEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Factory\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MainToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"concatenate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_percentRef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_opvAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniswapAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"joinEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"outEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refContract\",\"outputs\":[{\"internalType\":\"contractOPV_REF\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"setPercentRef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"}],\"name\":\"setTotalFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router02\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreateEventABI is the input ABI used to generate the binding from.
// Deprecated: Use CreateEventMetaData.ABI instead.
var CreateEventABI = CreateEventMetaData.ABI

// CreateEvent is an auto generated Go binding around an Ethereum contract.
type CreateEvent struct {
	CreateEventCaller     // Read-only binding to the contract
	CreateEventTransactor // Write-only binding to the contract
	CreateEventFilterer   // Log filterer for contract events
}

// CreateEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreateEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreateEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreateEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreateEventSession struct {
	Contract     *CreateEvent      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreateEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreateEventCallerSession struct {
	Contract *CreateEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CreateEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreateEventTransactorSession struct {
	Contract     *CreateEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CreateEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreateEventRaw struct {
	Contract *CreateEvent // Generic contract binding to access the raw methods on
}

// CreateEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreateEventCallerRaw struct {
	Contract *CreateEventCaller // Generic read-only contract binding to access the raw methods on
}

// CreateEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreateEventTransactorRaw struct {
	Contract *CreateEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreateEvent creates a new instance of CreateEvent, bound to a specific deployed contract.
func NewCreateEvent(address common.Address, backend bind.ContractBackend) (*CreateEvent, error) {
	contract, err := bindCreateEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreateEvent{CreateEventCaller: CreateEventCaller{contract: contract}, CreateEventTransactor: CreateEventTransactor{contract: contract}, CreateEventFilterer: CreateEventFilterer{contract: contract}}, nil
}

// NewCreateEventCaller creates a new read-only instance of CreateEvent, bound to a specific deployed contract.
func NewCreateEventCaller(address common.Address, caller bind.ContractCaller) (*CreateEventCaller, error) {
	contract, err := bindCreateEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreateEventCaller{contract: contract}, nil
}

// NewCreateEventTransactor creates a new write-only instance of CreateEvent, bound to a specific deployed contract.
func NewCreateEventTransactor(address common.Address, transactor bind.ContractTransactor) (*CreateEventTransactor, error) {
	contract, err := bindCreateEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreateEventTransactor{contract: contract}, nil
}

// NewCreateEventFilterer creates a new log filterer instance of CreateEvent, bound to a specific deployed contract.
func NewCreateEventFilterer(address common.Address, filterer bind.ContractFilterer) (*CreateEventFilterer, error) {
	contract, err := bindCreateEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreateEventFilterer{contract: contract}, nil
}

// bindCreateEvent binds a generic wrapper to an already deployed contract.
func bindCreateEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreateEventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateEvent *CreateEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreateEvent.Contract.CreateEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateEvent *CreateEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateEvent.Contract.CreateEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateEvent *CreateEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateEvent.Contract.CreateEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateEvent *CreateEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreateEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateEvent *CreateEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateEvent *CreateEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateEvent.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_CreateEvent *CreateEventCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_CreateEvent *CreateEventSession) Factory() (common.Address, error) {
	return _CreateEvent.Contract.Factory(&_CreateEvent.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc83dd231.
//
// Solidity: function Factory() view returns(address)
func (_CreateEvent *CreateEventCallerSession) Factory() (common.Address, error) {
	return _CreateEvent.Contract.Factory(&_CreateEvent.CallOpts)
}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_CreateEvent *CreateEventCaller) MainToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "MainToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_CreateEvent *CreateEventSession) MainToken() (common.Address, error) {
	return _CreateEvent.Contract.MainToken(&_CreateEvent.CallOpts)
}

// MainToken is a free data retrieval call binding the contract method 0xf5194ae2.
//
// Solidity: function MainToken() view returns(address)
func (_CreateEvent *CreateEventCallerSession) MainToken() (common.Address, error) {
	return _CreateEvent.Contract.MainToken(&_CreateEvent.CallOpts)
}

// CheckFee is a free data retrieval call binding the contract method 0x1904863c.
//
// Solidity: function checkFee() view returns(uint256)
func (_CreateEvent *CreateEventCaller) CheckFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "checkFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckFee is a free data retrieval call binding the contract method 0x1904863c.
//
// Solidity: function checkFee() view returns(uint256)
func (_CreateEvent *CreateEventSession) CheckFee() (*big.Int, error) {
	return _CreateEvent.Contract.CheckFee(&_CreateEvent.CallOpts)
}

// CheckFee is a free data retrieval call binding the contract method 0x1904863c.
//
// Solidity: function checkFee() view returns(uint256)
func (_CreateEvent *CreateEventCallerSession) CheckFee() (*big.Int, error) {
	return _CreateEvent.Contract.CheckFee(&_CreateEvent.CallOpts)
}

// Concatenate is a free data retrieval call binding the contract method 0x7d986470.
//
// Solidity: function concatenate(uint256 nftId, address contractAddress, address sender) pure returns(string)
func (_CreateEvent *CreateEventCaller) Concatenate(opts *bind.CallOpts, nftId *big.Int, contractAddress common.Address, sender common.Address) (string, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "concatenate", nftId, contractAddress, sender)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Concatenate is a free data retrieval call binding the contract method 0x7d986470.
//
// Solidity: function concatenate(uint256 nftId, address contractAddress, address sender) pure returns(string)
func (_CreateEvent *CreateEventSession) Concatenate(nftId *big.Int, contractAddress common.Address, sender common.Address) (string, error) {
	return _CreateEvent.Contract.Concatenate(&_CreateEvent.CallOpts, nftId, contractAddress, sender)
}

// Concatenate is a free data retrieval call binding the contract method 0x7d986470.
//
// Solidity: function concatenate(uint256 nftId, address contractAddress, address sender) pure returns(string)
func (_CreateEvent *CreateEventCallerSession) Concatenate(nftId *big.Int, contractAddress common.Address, sender common.Address) (string, error) {
	return _CreateEvent.Contract.Concatenate(&_CreateEvent.CallOpts, nftId, contractAddress, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreateEvent *CreateEventCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreateEvent *CreateEventSession) Owner() (common.Address, error) {
	return _CreateEvent.Contract.Owner(&_CreateEvent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreateEvent *CreateEventCallerSession) Owner() (common.Address, error) {
	return _CreateEvent.Contract.Owner(&_CreateEvent.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreateEvent *CreateEventCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreateEvent *CreateEventSession) Paused() (bool, error) {
	return _CreateEvent.Contract.Paused(&_CreateEvent.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreateEvent *CreateEventCallerSession) Paused() (bool, error) {
	return _CreateEvent.Contract.Paused(&_CreateEvent.CallOpts)
}

// RefContract is a free data retrieval call binding the contract method 0xdf97d31c.
//
// Solidity: function refContract() view returns(address)
func (_CreateEvent *CreateEventCaller) RefContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "refContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RefContract is a free data retrieval call binding the contract method 0xdf97d31c.
//
// Solidity: function refContract() view returns(address)
func (_CreateEvent *CreateEventSession) RefContract() (common.Address, error) {
	return _CreateEvent.Contract.RefContract(&_CreateEvent.CallOpts)
}

// RefContract is a free data retrieval call binding the contract method 0xdf97d31c.
//
// Solidity: function refContract() view returns(address)
func (_CreateEvent *CreateEventCallerSession) RefContract() (common.Address, error) {
	return _CreateEvent.Contract.RefContract(&_CreateEvent.CallOpts)
}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_CreateEvent *CreateEventCaller) UniswapV2Router02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreateEvent.contract.Call(opts, &out, "uniswapV2Router02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_CreateEvent *CreateEventSession) UniswapV2Router02() (common.Address, error) {
	return _CreateEvent.Contract.UniswapV2Router02(&_CreateEvent.CallOpts)
}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_CreateEvent *CreateEventCallerSession) UniswapV2Router02() (common.Address, error) {
	return _CreateEvent.Contract.UniswapV2Router02(&_CreateEvent.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xde17d3e3.
//
// Solidity: function init(uint256 _percentRef, uint256 _totalFee, address _feeAddress, address _opvAddress, address _refAddress, address _uniswapAddress, address _usdToken) returns()
func (_CreateEvent *CreateEventTransactor) Init(opts *bind.TransactOpts, _percentRef *big.Int, _totalFee *big.Int, _feeAddress common.Address, _opvAddress common.Address, _refAddress common.Address, _uniswapAddress common.Address, _usdToken common.Address) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "init", _percentRef, _totalFee, _feeAddress, _opvAddress, _refAddress, _uniswapAddress, _usdToken)
}

// Init is a paid mutator transaction binding the contract method 0xde17d3e3.
//
// Solidity: function init(uint256 _percentRef, uint256 _totalFee, address _feeAddress, address _opvAddress, address _refAddress, address _uniswapAddress, address _usdToken) returns()
func (_CreateEvent *CreateEventSession) Init(_percentRef *big.Int, _totalFee *big.Int, _feeAddress common.Address, _opvAddress common.Address, _refAddress common.Address, _uniswapAddress common.Address, _usdToken common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.Init(&_CreateEvent.TransactOpts, _percentRef, _totalFee, _feeAddress, _opvAddress, _refAddress, _uniswapAddress, _usdToken)
}

// Init is a paid mutator transaction binding the contract method 0xde17d3e3.
//
// Solidity: function init(uint256 _percentRef, uint256 _totalFee, address _feeAddress, address _opvAddress, address _refAddress, address _uniswapAddress, address _usdToken) returns()
func (_CreateEvent *CreateEventTransactorSession) Init(_percentRef *big.Int, _totalFee *big.Int, _feeAddress common.Address, _opvAddress common.Address, _refAddress common.Address, _uniswapAddress common.Address, _usdToken common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.Init(&_CreateEvent.TransactOpts, _percentRef, _totalFee, _feeAddress, _opvAddress, _refAddress, _uniswapAddress, _usdToken)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x3d0f2b53.
//
// Solidity: function joinEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventTransactor) JoinEvent(opts *bind.TransactOpts, nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "joinEvent", nftId, contractAddress)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x3d0f2b53.
//
// Solidity: function joinEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventSession) JoinEvent(nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.JoinEvent(&_CreateEvent.TransactOpts, nftId, contractAddress)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x3d0f2b53.
//
// Solidity: function joinEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventTransactorSession) JoinEvent(nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.JoinEvent(&_CreateEvent.TransactOpts, nftId, contractAddress)
}

// OutEvent is a paid mutator transaction binding the contract method 0x58e97d6a.
//
// Solidity: function outEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventTransactor) OutEvent(opts *bind.TransactOpts, nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "outEvent", nftId, contractAddress)
}

// OutEvent is a paid mutator transaction binding the contract method 0x58e97d6a.
//
// Solidity: function outEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventSession) OutEvent(nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.OutEvent(&_CreateEvent.TransactOpts, nftId, contractAddress)
}

// OutEvent is a paid mutator transaction binding the contract method 0x58e97d6a.
//
// Solidity: function outEvent(uint256 nftId, address contractAddress) returns()
func (_CreateEvent *CreateEventTransactorSession) OutEvent(nftId *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.OutEvent(&_CreateEvent.TransactOpts, nftId, contractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreateEvent *CreateEventTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreateEvent *CreateEventSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreateEvent.Contract.RenounceOwnership(&_CreateEvent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreateEvent *CreateEventTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreateEvent.Contract.RenounceOwnership(&_CreateEvent.TransactOpts)
}

// SetPercentRef is a paid mutator transaction binding the contract method 0xa9a4e672.
//
// Solidity: function setPercentRef(uint256 _newPercent) returns()
func (_CreateEvent *CreateEventTransactor) SetPercentRef(opts *bind.TransactOpts, _newPercent *big.Int) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "setPercentRef", _newPercent)
}

// SetPercentRef is a paid mutator transaction binding the contract method 0xa9a4e672.
//
// Solidity: function setPercentRef(uint256 _newPercent) returns()
func (_CreateEvent *CreateEventSession) SetPercentRef(_newPercent *big.Int) (*types.Transaction, error) {
	return _CreateEvent.Contract.SetPercentRef(&_CreateEvent.TransactOpts, _newPercent)
}

// SetPercentRef is a paid mutator transaction binding the contract method 0xa9a4e672.
//
// Solidity: function setPercentRef(uint256 _newPercent) returns()
func (_CreateEvent *CreateEventTransactorSession) SetPercentRef(_newPercent *big.Int) (*types.Transaction, error) {
	return _CreateEvent.Contract.SetPercentRef(&_CreateEvent.TransactOpts, _newPercent)
}

// SetTotalFee is a paid mutator transaction binding the contract method 0x992d0ebb.
//
// Solidity: function setTotalFee(uint256 _totalFee) returns()
func (_CreateEvent *CreateEventTransactor) SetTotalFee(opts *bind.TransactOpts, _totalFee *big.Int) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "setTotalFee", _totalFee)
}

// SetTotalFee is a paid mutator transaction binding the contract method 0x992d0ebb.
//
// Solidity: function setTotalFee(uint256 _totalFee) returns()
func (_CreateEvent *CreateEventSession) SetTotalFee(_totalFee *big.Int) (*types.Transaction, error) {
	return _CreateEvent.Contract.SetTotalFee(&_CreateEvent.TransactOpts, _totalFee)
}

// SetTotalFee is a paid mutator transaction binding the contract method 0x992d0ebb.
//
// Solidity: function setTotalFee(uint256 _totalFee) returns()
func (_CreateEvent *CreateEventTransactorSession) SetTotalFee(_totalFee *big.Int) (*types.Transaction, error) {
	return _CreateEvent.Contract.SetTotalFee(&_CreateEvent.TransactOpts, _totalFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreateEvent *CreateEventTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreateEvent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreateEvent *CreateEventSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.TransferOwnership(&_CreateEvent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreateEvent *CreateEventTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreateEvent.Contract.TransferOwnership(&_CreateEvent.TransactOpts, newOwner)
}

// CreateEventJoinEventIterator is returned from FilterJoinEvent and is used to iterate over the raw logs and unpacked data for JoinEvent events raised by the CreateEvent contract.
type CreateEventJoinEventIterator struct {
	Event *CreateEventJoinEvent // Event containing the contract specifics and raw log

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
func (it *CreateEventJoinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreateEventJoinEvent)
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
		it.Event = new(CreateEventJoinEvent)
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
func (it *CreateEventJoinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreateEventJoinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreateEventJoinEvent represents a JoinEvent event raised by the CreateEvent contract.
type CreateEventJoinEvent struct {
	NftId           *big.Int
	ContractAddress common.Address
	User            common.Address
	FundToSystem    *big.Int
	FundToRef       *big.Int
	TotalFund       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterJoinEvent is a free log retrieval operation binding the contract event 0x0780f4e74c6e9d83a59cc582651e73595827608dd659f675debaed1e691f0945.
//
// Solidity: event JoinEvent(uint256 nftId, address contractAddress, address user, uint256 fundToSystem, uint256 fundToRef, uint256 totalFund)
func (_CreateEvent *CreateEventFilterer) FilterJoinEvent(opts *bind.FilterOpts) (*CreateEventJoinEventIterator, error) {

	logs, sub, err := _CreateEvent.contract.FilterLogs(opts, "JoinEvent")
	if err != nil {
		return nil, err
	}
	return &CreateEventJoinEventIterator{contract: _CreateEvent.contract, event: "JoinEvent", logs: logs, sub: sub}, nil
}

// WatchJoinEvent is a free log subscription operation binding the contract event 0x0780f4e74c6e9d83a59cc582651e73595827608dd659f675debaed1e691f0945.
//
// Solidity: event JoinEvent(uint256 nftId, address contractAddress, address user, uint256 fundToSystem, uint256 fundToRef, uint256 totalFund)
func (_CreateEvent *CreateEventFilterer) WatchJoinEvent(opts *bind.WatchOpts, sink chan<- *CreateEventJoinEvent) (event.Subscription, error) {

	logs, sub, err := _CreateEvent.contract.WatchLogs(opts, "JoinEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreateEventJoinEvent)
				if err := _CreateEvent.contract.UnpackLog(event, "JoinEvent", log); err != nil {
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

// ParseJoinEvent is a log parse operation binding the contract event 0x0780f4e74c6e9d83a59cc582651e73595827608dd659f675debaed1e691f0945.
//
// Solidity: event JoinEvent(uint256 nftId, address contractAddress, address user, uint256 fundToSystem, uint256 fundToRef, uint256 totalFund)
func (_CreateEvent *CreateEventFilterer) ParseJoinEvent(log types.Log) (*CreateEventJoinEvent, error) {
	event := new(CreateEventJoinEvent)
	if err := _CreateEvent.contract.UnpackLog(event, "JoinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreateEventOutEventIterator is returned from FilterOutEvent and is used to iterate over the raw logs and unpacked data for OutEvent events raised by the CreateEvent contract.
type CreateEventOutEventIterator struct {
	Event *CreateEventOutEvent // Event containing the contract specifics and raw log

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
func (it *CreateEventOutEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreateEventOutEvent)
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
		it.Event = new(CreateEventOutEvent)
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
func (it *CreateEventOutEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreateEventOutEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreateEventOutEvent represents a OutEvent event raised by the CreateEvent contract.
type CreateEventOutEvent struct {
	NftId           *big.Int
	ContractAddress common.Address
	User            common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOutEvent is a free log retrieval operation binding the contract event 0x4ebb077117e5db88e5cda8ca72c224ac74ee9ac1b182f2417853f4decbfdc476.
//
// Solidity: event OutEvent(uint256 nftId, address contractAddress, address user)
func (_CreateEvent *CreateEventFilterer) FilterOutEvent(opts *bind.FilterOpts) (*CreateEventOutEventIterator, error) {

	logs, sub, err := _CreateEvent.contract.FilterLogs(opts, "OutEvent")
	if err != nil {
		return nil, err
	}
	return &CreateEventOutEventIterator{contract: _CreateEvent.contract, event: "OutEvent", logs: logs, sub: sub}, nil
}

// WatchOutEvent is a free log subscription operation binding the contract event 0x4ebb077117e5db88e5cda8ca72c224ac74ee9ac1b182f2417853f4decbfdc476.
//
// Solidity: event OutEvent(uint256 nftId, address contractAddress, address user)
func (_CreateEvent *CreateEventFilterer) WatchOutEvent(opts *bind.WatchOpts, sink chan<- *CreateEventOutEvent) (event.Subscription, error) {

	logs, sub, err := _CreateEvent.contract.WatchLogs(opts, "OutEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreateEventOutEvent)
				if err := _CreateEvent.contract.UnpackLog(event, "OutEvent", log); err != nil {
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

// ParseOutEvent is a log parse operation binding the contract event 0x4ebb077117e5db88e5cda8ca72c224ac74ee9ac1b182f2417853f4decbfdc476.
//
// Solidity: event OutEvent(uint256 nftId, address contractAddress, address user)
func (_CreateEvent *CreateEventFilterer) ParseOutEvent(log types.Log) (*CreateEventOutEvent, error) {
	event := new(CreateEventOutEvent)
	if err := _CreateEvent.contract.UnpackLog(event, "OutEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreateEventOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CreateEvent contract.
type CreateEventOwnershipTransferredIterator struct {
	Event *CreateEventOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreateEventOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreateEventOwnershipTransferred)
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
		it.Event = new(CreateEventOwnershipTransferred)
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
func (it *CreateEventOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreateEventOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreateEventOwnershipTransferred represents a OwnershipTransferred event raised by the CreateEvent contract.
type CreateEventOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreateEvent *CreateEventFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreateEventOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreateEvent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreateEventOwnershipTransferredIterator{contract: _CreateEvent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreateEvent *CreateEventFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreateEventOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreateEvent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreateEventOwnershipTransferred)
				if err := _CreateEvent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreateEvent *CreateEventFilterer) ParseOwnershipTransferred(log types.Log) (*CreateEventOwnershipTransferred, error) {
	event := new(CreateEventOwnershipTransferred)
	if err := _CreateEvent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreateEventPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreateEvent contract.
type CreateEventPausedIterator struct {
	Event *CreateEventPaused // Event containing the contract specifics and raw log

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
func (it *CreateEventPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreateEventPaused)
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
		it.Event = new(CreateEventPaused)
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
func (it *CreateEventPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreateEventPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreateEventPaused represents a Paused event raised by the CreateEvent contract.
type CreateEventPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreateEvent *CreateEventFilterer) FilterPaused(opts *bind.FilterOpts) (*CreateEventPausedIterator, error) {

	logs, sub, err := _CreateEvent.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreateEventPausedIterator{contract: _CreateEvent.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreateEvent *CreateEventFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreateEventPaused) (event.Subscription, error) {

	logs, sub, err := _CreateEvent.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreateEventPaused)
				if err := _CreateEvent.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreateEvent *CreateEventFilterer) ParsePaused(log types.Log) (*CreateEventPaused, error) {
	event := new(CreateEventPaused)
	if err := _CreateEvent.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreateEventUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreateEvent contract.
type CreateEventUnpausedIterator struct {
	Event *CreateEventUnpaused // Event containing the contract specifics and raw log

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
func (it *CreateEventUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreateEventUnpaused)
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
		it.Event = new(CreateEventUnpaused)
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
func (it *CreateEventUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreateEventUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreateEventUnpaused represents a Unpaused event raised by the CreateEvent contract.
type CreateEventUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreateEvent *CreateEventFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreateEventUnpausedIterator, error) {

	logs, sub, err := _CreateEvent.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreateEventUnpausedIterator{contract: _CreateEvent.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreateEvent *CreateEventFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreateEventUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreateEvent.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreateEventUnpaused)
				if err := _CreateEvent.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreateEvent *CreateEventFilterer) ParseUnpaused(log types.Log) (*CreateEventUnpaused, error) {
	event := new(CreateEventUnpaused)
	if err := _CreateEvent.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
