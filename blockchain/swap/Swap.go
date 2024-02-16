// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// OpvSwapMetaData contains all meta data concerning the OpvSwap contract.
var OpvSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uniswap_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stable_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_opv\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenSource\",\"type\":\"address\"}],\"name\":\"getEst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenSource\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router02\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OpvSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use OpvSwapMetaData.ABI instead.
var OpvSwapABI = OpvSwapMetaData.ABI

// OpvSwap is an auto generated Go binding around an Ethereum contract.
type OpvSwap struct {
	OpvSwapCaller     // Read-only binding to the contract
	OpvSwapTransactor // Write-only binding to the contract
	OpvSwapFilterer   // Log filterer for contract events
}

// OpvSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpvSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpvSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpvSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpvSwapSession struct {
	Contract     *OpvSwap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpvSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpvSwapCallerSession struct {
	Contract *OpvSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpvSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpvSwapTransactorSession struct {
	Contract     *OpvSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpvSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpvSwapRaw struct {
	Contract *OpvSwap // Generic contract binding to access the raw methods on
}

// OpvSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpvSwapCallerRaw struct {
	Contract *OpvSwapCaller // Generic read-only contract binding to access the raw methods on
}

// OpvSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpvSwapTransactorRaw struct {
	Contract *OpvSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpvSwap creates a new instance of OpvSwap, bound to a specific deployed contract.
func NewOpvSwap(address common.Address, backend bind.ContractBackend) (*OpvSwap, error) {
	contract, err := bindOpvSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpvSwap{OpvSwapCaller: OpvSwapCaller{contract: contract}, OpvSwapTransactor: OpvSwapTransactor{contract: contract}, OpvSwapFilterer: OpvSwapFilterer{contract: contract}}, nil
}

// NewOpvSwapCaller creates a new read-only instance of OpvSwap, bound to a specific deployed contract.
func NewOpvSwapCaller(address common.Address, caller bind.ContractCaller) (*OpvSwapCaller, error) {
	contract, err := bindOpvSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpvSwapCaller{contract: contract}, nil
}

// NewOpvSwapTransactor creates a new write-only instance of OpvSwap, bound to a specific deployed contract.
func NewOpvSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*OpvSwapTransactor, error) {
	contract, err := bindOpvSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpvSwapTransactor{contract: contract}, nil
}

// NewOpvSwapFilterer creates a new log filterer instance of OpvSwap, bound to a specific deployed contract.
func NewOpvSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*OpvSwapFilterer, error) {
	contract, err := bindOpvSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpvSwapFilterer{contract: contract}, nil
}

// bindOpvSwap binds a generic wrapper to an already deployed contract.
func bindOpvSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpvSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvSwap *OpvSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvSwap.Contract.OpvSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvSwap *OpvSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvSwap.Contract.OpvSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvSwap *OpvSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvSwap.Contract.OpvSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvSwap *OpvSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvSwap *OpvSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvSwap *OpvSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvSwap.Contract.contract.Transact(opts, method, params...)
}

// GetEst is a free data retrieval call binding the contract method 0x1b2393c3.
//
// Solidity: function getEst(uint256 _amountTokenIn, address _tokenSource) view returns(uint256)
func (_OpvSwap *OpvSwapCaller) GetEst(opts *bind.CallOpts, _amountTokenIn *big.Int, _tokenSource common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OpvSwap.contract.Call(opts, &out, "getEst", _amountTokenIn, _tokenSource)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEst is a free data retrieval call binding the contract method 0x1b2393c3.
//
// Solidity: function getEst(uint256 _amountTokenIn, address _tokenSource) view returns(uint256)
func (_OpvSwap *OpvSwapSession) GetEst(_amountTokenIn *big.Int, _tokenSource common.Address) (*big.Int, error) {
	return _OpvSwap.Contract.GetEst(&_OpvSwap.CallOpts, _amountTokenIn, _tokenSource)
}

// GetEst is a free data retrieval call binding the contract method 0x1b2393c3.
//
// Solidity: function getEst(uint256 _amountTokenIn, address _tokenSource) view returns(uint256)
func (_OpvSwap *OpvSwapCallerSession) GetEst(_amountTokenIn *big.Int, _tokenSource common.Address) (*big.Int, error) {
	return _OpvSwap.Contract.GetEst(&_OpvSwap.CallOpts, _amountTokenIn, _tokenSource)
}

// Opv is a free data retrieval call binding the contract method 0xf2481fca.
//
// Solidity: function opv() view returns(address)
func (_OpvSwap *OpvSwapCaller) Opv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvSwap.contract.Call(opts, &out, "opv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Opv is a free data retrieval call binding the contract method 0xf2481fca.
//
// Solidity: function opv() view returns(address)
func (_OpvSwap *OpvSwapSession) Opv() (common.Address, error) {
	return _OpvSwap.Contract.Opv(&_OpvSwap.CallOpts)
}

// Opv is a free data retrieval call binding the contract method 0xf2481fca.
//
// Solidity: function opv() view returns(address)
func (_OpvSwap *OpvSwapCallerSession) Opv() (common.Address, error) {
	return _OpvSwap.Contract.Opv(&_OpvSwap.CallOpts)
}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_OpvSwap *OpvSwapCaller) StableToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvSwap.contract.Call(opts, &out, "stableToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_OpvSwap *OpvSwapSession) StableToken() (common.Address, error) {
	return _OpvSwap.Contract.StableToken(&_OpvSwap.CallOpts)
}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_OpvSwap *OpvSwapCallerSession) StableToken() (common.Address, error) {
	return _OpvSwap.Contract.StableToken(&_OpvSwap.CallOpts)
}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_OpvSwap *OpvSwapCaller) UniswapV2Router02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvSwap.contract.Call(opts, &out, "uniswapV2Router02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_OpvSwap *OpvSwapSession) UniswapV2Router02() (common.Address, error) {
	return _OpvSwap.Contract.UniswapV2Router02(&_OpvSwap.CallOpts)
}

// UniswapV2Router02 is a free data retrieval call binding the contract method 0xa7c6402c.
//
// Solidity: function uniswapV2Router02() view returns(address)
func (_OpvSwap *OpvSwapCallerSession) UniswapV2Router02() (common.Address, error) {
	return _OpvSwap.Contract.UniswapV2Router02(&_OpvSwap.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xd3986f08.
//
// Solidity: function swap(uint256 _amountTokenIn, address _tokenSource) returns()
func (_OpvSwap *OpvSwapTransactor) Swap(opts *bind.TransactOpts, _amountTokenIn *big.Int, _tokenSource common.Address) (*types.Transaction, error) {
	return _OpvSwap.contract.Transact(opts, "swap", _amountTokenIn, _tokenSource)
}

// Swap is a paid mutator transaction binding the contract method 0xd3986f08.
//
// Solidity: function swap(uint256 _amountTokenIn, address _tokenSource) returns()
func (_OpvSwap *OpvSwapSession) Swap(_amountTokenIn *big.Int, _tokenSource common.Address) (*types.Transaction, error) {
	return _OpvSwap.Contract.Swap(&_OpvSwap.TransactOpts, _amountTokenIn, _tokenSource)
}

// Swap is a paid mutator transaction binding the contract method 0xd3986f08.
//
// Solidity: function swap(uint256 _amountTokenIn, address _tokenSource) returns()
func (_OpvSwap *OpvSwapTransactorSession) Swap(_amountTokenIn *big.Int, _tokenSource common.Address) (*types.Transaction, error) {
	return _OpvSwap.Contract.Swap(&_OpvSwap.TransactOpts, _amountTokenIn, _tokenSource)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OpvSwap *OpvSwapTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OpvSwap.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OpvSwap *OpvSwapSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OpvSwap.Contract.Fallback(&_OpvSwap.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OpvSwap *OpvSwapTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OpvSwap.Contract.Fallback(&_OpvSwap.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvSwap *OpvSwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvSwap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvSwap *OpvSwapSession) Receive() (*types.Transaction, error) {
	return _OpvSwap.Contract.Receive(&_OpvSwap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvSwap *OpvSwapTransactorSession) Receive() (*types.Transaction, error) {
	return _OpvSwap.Contract.Receive(&_OpvSwap.TransactOpts)
}
