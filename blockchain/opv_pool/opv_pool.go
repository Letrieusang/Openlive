// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opv_pool

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

// OpvPoolMetaData contains all meta data concerning the OpvPool contract.
var OpvPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PERCENTS_DIVIDER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeTransfer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_numTokensSellToAddToLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"antiBotEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTransfer\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeWallet\",\"type\":\"address\"}],\"name\":\"changeFeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromBot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedToFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newAddr\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAddr\",\"type\":\"address[]\"}],\"name\":\"modifyBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newAddr\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAddr\",\"type\":\"address[]\"}],\"name\":\"modifyWhiteListBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newAddr\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAddr\",\"type\":\"address[]\"}],\"name\":\"modifyWhiteListPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newAddr\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAddr\",\"type\":\"address[]\"}],\"name\":\"modifyWhiteListReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newAddr\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedAddr\",\"type\":\"address[]\"}],\"name\":\"modifyWhiteListSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setAntiBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setLimitSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokensSellToAddToLiquidity\",\"type\":\"uint256\"}],\"name\":\"setNumTokensSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setSwapWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListBot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListReceiver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OpvPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use OpvPoolMetaData.ABI instead.
var OpvPoolABI = OpvPoolMetaData.ABI

// OpvPool is an auto generated Go binding around an Ethereum contract.
type OpvPool struct {
	OpvPoolCaller     // Read-only binding to the contract
	OpvPoolTransactor // Write-only binding to the contract
	OpvPoolFilterer   // Log filterer for contract events
}

// OpvPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpvPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpvPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpvPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpvPoolSession struct {
	Contract     *OpvPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpvPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpvPoolCallerSession struct {
	Contract *OpvPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpvPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpvPoolTransactorSession struct {
	Contract     *OpvPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpvPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpvPoolRaw struct {
	Contract *OpvPool // Generic contract binding to access the raw methods on
}

// OpvPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpvPoolCallerRaw struct {
	Contract *OpvPoolCaller // Generic read-only contract binding to access the raw methods on
}

// OpvPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpvPoolTransactorRaw struct {
	Contract *OpvPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpvPool creates a new instance of OpvPool, bound to a specific deployed contract.
func NewOpvPool(address common.Address, backend bind.ContractBackend) (*OpvPool, error) {
	contract, err := bindOpvPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpvPool{OpvPoolCaller: OpvPoolCaller{contract: contract}, OpvPoolTransactor: OpvPoolTransactor{contract: contract}, OpvPoolFilterer: OpvPoolFilterer{contract: contract}}, nil
}

// NewOpvPoolCaller creates a new read-only instance of OpvPool, bound to a specific deployed contract.
func NewOpvPoolCaller(address common.Address, caller bind.ContractCaller) (*OpvPoolCaller, error) {
	contract, err := bindOpvPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpvPoolCaller{contract: contract}, nil
}

// NewOpvPoolTransactor creates a new write-only instance of OpvPool, bound to a specific deployed contract.
func NewOpvPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*OpvPoolTransactor, error) {
	contract, err := bindOpvPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpvPoolTransactor{contract: contract}, nil
}

// NewOpvPoolFilterer creates a new log filterer instance of OpvPool, bound to a specific deployed contract.
func NewOpvPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*OpvPoolFilterer, error) {
	contract, err := bindOpvPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpvPoolFilterer{contract: contract}, nil
}

// bindOpvPool binds a generic wrapper to an already deployed contract.
func bindOpvPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpvPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvPool *OpvPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvPool.Contract.OpvPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvPool *OpvPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvPool.Contract.OpvPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvPool *OpvPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvPool.Contract.OpvPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvPool *OpvPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvPool *OpvPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvPool *OpvPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvPool.Contract.contract.Transact(opts, method, params...)
}

// PERCENTSDIVIDER is a free data retrieval call binding the contract method 0x01c234a8.
//
// Solidity: function PERCENTS_DIVIDER() view returns(uint256)
func (_OpvPool *OpvPoolCaller) PERCENTSDIVIDER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "PERCENTS_DIVIDER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTSDIVIDER is a free data retrieval call binding the contract method 0x01c234a8.
//
// Solidity: function PERCENTS_DIVIDER() view returns(uint256)
func (_OpvPool *OpvPoolSession) PERCENTSDIVIDER() (*big.Int, error) {
	return _OpvPool.Contract.PERCENTSDIVIDER(&_OpvPool.CallOpts)
}

// PERCENTSDIVIDER is a free data retrieval call binding the contract method 0x01c234a8.
//
// Solidity: function PERCENTS_DIVIDER() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) PERCENTSDIVIDER() (*big.Int, error) {
	return _OpvPool.Contract.PERCENTSDIVIDER(&_OpvPool.CallOpts)
}

// FeeTransfer is a free data retrieval call binding the contract method 0x0b38dd3e.
//
// Solidity: function _feeTransfer() view returns(uint256)
func (_OpvPool *OpvPoolCaller) FeeTransfer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "_feeTransfer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeTransfer is a free data retrieval call binding the contract method 0x0b38dd3e.
//
// Solidity: function _feeTransfer() view returns(uint256)
func (_OpvPool *OpvPoolSession) FeeTransfer() (*big.Int, error) {
	return _OpvPool.Contract.FeeTransfer(&_OpvPool.CallOpts)
}

// FeeTransfer is a free data retrieval call binding the contract method 0x0b38dd3e.
//
// Solidity: function _feeTransfer() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) FeeTransfer() (*big.Int, error) {
	return _OpvPool.Contract.FeeTransfer(&_OpvPool.CallOpts)
}

// FeeWallet is a free data retrieval call binding the contract method 0x659419a4.
//
// Solidity: function _feeWallet() view returns(address)
func (_OpvPool *OpvPoolCaller) FeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "_feeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeWallet is a free data retrieval call binding the contract method 0x659419a4.
//
// Solidity: function _feeWallet() view returns(address)
func (_OpvPool *OpvPoolSession) FeeWallet() (common.Address, error) {
	return _OpvPool.Contract.FeeWallet(&_OpvPool.CallOpts)
}

// FeeWallet is a free data retrieval call binding the contract method 0x659419a4.
//
// Solidity: function _feeWallet() view returns(address)
func (_OpvPool *OpvPoolCallerSession) FeeWallet() (common.Address, error) {
	return _OpvPool.Contract.FeeWallet(&_OpvPool.CallOpts)
}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xb3f22ce3.
//
// Solidity: function _numTokensSellToAddToLiquidity() view returns(uint256)
func (_OpvPool *OpvPoolCaller) NumTokensSellToAddToLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "_numTokensSellToAddToLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xb3f22ce3.
//
// Solidity: function _numTokensSellToAddToLiquidity() view returns(uint256)
func (_OpvPool *OpvPoolSession) NumTokensSellToAddToLiquidity() (*big.Int, error) {
	return _OpvPool.Contract.NumTokensSellToAddToLiquidity(&_OpvPool.CallOpts)
}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xb3f22ce3.
//
// Solidity: function _numTokensSellToAddToLiquidity() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) NumTokensSellToAddToLiquidity() (*big.Int, error) {
	return _OpvPool.Contract.NumTokensSellToAddToLiquidity(&_OpvPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OpvPool *OpvPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OpvPool *OpvPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OpvPool.Contract.Allowance(&_OpvPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OpvPool.Contract.Allowance(&_OpvPool.CallOpts, owner, spender)
}

// AntiBotEnabled is a free data retrieval call binding the contract method 0xd8c6404b.
//
// Solidity: function antiBotEnabled() view returns(bool)
func (_OpvPool *OpvPoolCaller) AntiBotEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "antiBotEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AntiBotEnabled is a free data retrieval call binding the contract method 0xd8c6404b.
//
// Solidity: function antiBotEnabled() view returns(bool)
func (_OpvPool *OpvPoolSession) AntiBotEnabled() (bool, error) {
	return _OpvPool.Contract.AntiBotEnabled(&_OpvPool.CallOpts)
}

// AntiBotEnabled is a free data retrieval call binding the contract method 0xd8c6404b.
//
// Solidity: function antiBotEnabled() view returns(bool)
func (_OpvPool *OpvPoolCallerSession) AntiBotEnabled() (bool, error) {
	return _OpvPool.Contract.AntiBotEnabled(&_OpvPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OpvPool *OpvPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OpvPool *OpvPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OpvPool.Contract.BalanceOf(&_OpvPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OpvPool.Contract.BalanceOf(&_OpvPool.CallOpts, account)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_OpvPool *OpvPoolCaller) BlackList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "blackList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_OpvPool *OpvPoolSession) BlackList(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.BlackList(&_OpvPool.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) BlackList(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.BlackList(&_OpvPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OpvPool *OpvPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OpvPool *OpvPoolSession) Decimals() (uint8, error) {
	return _OpvPool.Contract.Decimals(&_OpvPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OpvPool *OpvPoolCallerSession) Decimals() (uint8, error) {
	return _OpvPool.Contract.Decimals(&_OpvPool.CallOpts)
}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_OpvPool *OpvPoolCaller) GeUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "geUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_OpvPool *OpvPoolSession) GeUnlockTime() (*big.Int, error) {
	return _OpvPool.Contract.GeUnlockTime(&_OpvPool.CallOpts)
}

// GeUnlockTime is a free data retrieval call binding the contract method 0xb6c52324.
//
// Solidity: function geUnlockTime() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) GeUnlockTime() (*big.Int, error) {
	return _OpvPool.Contract.GeUnlockTime(&_OpvPool.CallOpts)
}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address account) view returns(bool)
func (_OpvPool *OpvPoolCaller) IsBlackList(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "isBlackList", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address account) view returns(bool)
func (_OpvPool *OpvPoolSession) IsBlackList(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsBlackList(&_OpvPool.CallOpts, account)
}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address account) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) IsBlackList(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsBlackList(&_OpvPool.CallOpts, account)
}

// IsExcludedFromBot is a free data retrieval call binding the contract method 0x1f6d1d75.
//
// Solidity: function isExcludedFromBot(address account) view returns(bool)
func (_OpvPool *OpvPoolCaller) IsExcludedFromBot(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "isExcludedFromBot", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromBot is a free data retrieval call binding the contract method 0x1f6d1d75.
//
// Solidity: function isExcludedFromBot(address account) view returns(bool)
func (_OpvPool *OpvPoolSession) IsExcludedFromBot(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromBot(&_OpvPool.CallOpts, account)
}

// IsExcludedFromBot is a free data retrieval call binding the contract method 0x1f6d1d75.
//
// Solidity: function isExcludedFromBot(address account) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) IsExcludedFromBot(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromBot(&_OpvPool.CallOpts, account)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_OpvPool *OpvPoolCaller) IsExcludedFromFee(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "isExcludedFromFee", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_OpvPool *OpvPoolSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromFee(&_OpvPool.CallOpts, account)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromFee(&_OpvPool.CallOpts, account)
}

// IsExcludedFromPool is a free data retrieval call binding the contract method 0x32a88ea2.
//
// Solidity: function isExcludedFromPool(address account) view returns(bool)
func (_OpvPool *OpvPoolCaller) IsExcludedFromPool(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "isExcludedFromPool", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromPool is a free data retrieval call binding the contract method 0x32a88ea2.
//
// Solidity: function isExcludedFromPool(address account) view returns(bool)
func (_OpvPool *OpvPoolSession) IsExcludedFromPool(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromPool(&_OpvPool.CallOpts, account)
}

// IsExcludedFromPool is a free data retrieval call binding the contract method 0x32a88ea2.
//
// Solidity: function isExcludedFromPool(address account) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) IsExcludedFromPool(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedFromPool(&_OpvPool.CallOpts, account)
}

// IsExcludedToFee is a free data retrieval call binding the contract method 0x844f30fb.
//
// Solidity: function isExcludedToFee(address account) view returns(bool)
func (_OpvPool *OpvPoolCaller) IsExcludedToFee(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "isExcludedToFee", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedToFee is a free data retrieval call binding the contract method 0x844f30fb.
//
// Solidity: function isExcludedToFee(address account) view returns(bool)
func (_OpvPool *OpvPoolSession) IsExcludedToFee(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedToFee(&_OpvPool.CallOpts, account)
}

// IsExcludedToFee is a free data retrieval call binding the contract method 0x844f30fb.
//
// Solidity: function isExcludedToFee(address account) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) IsExcludedToFee(account common.Address) (bool, error) {
	return _OpvPool.Contract.IsExcludedToFee(&_OpvPool.CallOpts, account)
}

// LimitSell is a free data retrieval call binding the contract method 0xc8972028.
//
// Solidity: function limitSell() view returns(bool)
func (_OpvPool *OpvPoolCaller) LimitSell(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "limitSell")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LimitSell is a free data retrieval call binding the contract method 0xc8972028.
//
// Solidity: function limitSell() view returns(bool)
func (_OpvPool *OpvPoolSession) LimitSell() (bool, error) {
	return _OpvPool.Contract.LimitSell(&_OpvPool.CallOpts)
}

// LimitSell is a free data retrieval call binding the contract method 0xc8972028.
//
// Solidity: function limitSell() view returns(bool)
func (_OpvPool *OpvPoolCallerSession) LimitSell() (bool, error) {
	return _OpvPool.Contract.LimitSell(&_OpvPool.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_OpvPool *OpvPoolCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_OpvPool *OpvPoolSession) MaxSupply() (*big.Int, error) {
	return _OpvPool.Contract.MaxSupply(&_OpvPool.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) MaxSupply() (*big.Int, error) {
	return _OpvPool.Contract.MaxSupply(&_OpvPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvPool *OpvPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvPool *OpvPoolSession) Name() (string, error) {
	return _OpvPool.Contract.Name(&_OpvPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvPool *OpvPoolCallerSession) Name() (string, error) {
	return _OpvPool.Contract.Name(&_OpvPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvPool *OpvPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvPool *OpvPoolSession) Owner() (common.Address, error) {
	return _OpvPool.Contract.Owner(&_OpvPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvPool *OpvPoolCallerSession) Owner() (common.Address, error) {
	return _OpvPool.Contract.Owner(&_OpvPool.CallOpts)
}

// SwapWhiteList is a free data retrieval call binding the contract method 0x31428a41.
//
// Solidity: function swapWhiteList() view returns(bool)
func (_OpvPool *OpvPoolCaller) SwapWhiteList(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "swapWhiteList")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapWhiteList is a free data retrieval call binding the contract method 0x31428a41.
//
// Solidity: function swapWhiteList() view returns(bool)
func (_OpvPool *OpvPoolSession) SwapWhiteList() (bool, error) {
	return _OpvPool.Contract.SwapWhiteList(&_OpvPool.CallOpts)
}

// SwapWhiteList is a free data retrieval call binding the contract method 0x31428a41.
//
// Solidity: function swapWhiteList() view returns(bool)
func (_OpvPool *OpvPoolCallerSession) SwapWhiteList() (bool, error) {
	return _OpvPool.Contract.SwapWhiteList(&_OpvPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvPool *OpvPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvPool *OpvPoolSession) Symbol() (string, error) {
	return _OpvPool.Contract.Symbol(&_OpvPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvPool *OpvPoolCallerSession) Symbol() (string, error) {
	return _OpvPool.Contract.Symbol(&_OpvPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvPool *OpvPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvPool *OpvPoolSession) TotalSupply() (*big.Int, error) {
	return _OpvPool.Contract.TotalSupply(&_OpvPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvPool *OpvPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _OpvPool.Contract.TotalSupply(&_OpvPool.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_OpvPool *OpvPoolCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_OpvPool *OpvPoolSession) UniswapV2Pair() (common.Address, error) {
	return _OpvPool.Contract.UniswapV2Pair(&_OpvPool.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_OpvPool *OpvPoolCallerSession) UniswapV2Pair() (common.Address, error) {
	return _OpvPool.Contract.UniswapV2Pair(&_OpvPool.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_OpvPool *OpvPoolCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_OpvPool *OpvPoolSession) UniswapV2Router() (common.Address, error) {
	return _OpvPool.Contract.UniswapV2Router(&_OpvPool.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_OpvPool *OpvPoolCallerSession) UniswapV2Router() (common.Address, error) {
	return _OpvPool.Contract.UniswapV2Router(&_OpvPool.CallOpts)
}

// WhiteListBot is a free data retrieval call binding the contract method 0x643b0de4.
//
// Solidity: function whiteListBot(address ) view returns(bool)
func (_OpvPool *OpvPoolCaller) WhiteListBot(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "whiteListBot", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListBot is a free data retrieval call binding the contract method 0x643b0de4.
//
// Solidity: function whiteListBot(address ) view returns(bool)
func (_OpvPool *OpvPoolSession) WhiteListBot(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListBot(&_OpvPool.CallOpts, arg0)
}

// WhiteListBot is a free data retrieval call binding the contract method 0x643b0de4.
//
// Solidity: function whiteListBot(address ) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) WhiteListBot(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListBot(&_OpvPool.CallOpts, arg0)
}

// WhiteListPool is a free data retrieval call binding the contract method 0x0f61b70e.
//
// Solidity: function whiteListPool(address ) view returns(bool)
func (_OpvPool *OpvPoolCaller) WhiteListPool(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "whiteListPool", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListPool is a free data retrieval call binding the contract method 0x0f61b70e.
//
// Solidity: function whiteListPool(address ) view returns(bool)
func (_OpvPool *OpvPoolSession) WhiteListPool(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListPool(&_OpvPool.CallOpts, arg0)
}

// WhiteListPool is a free data retrieval call binding the contract method 0x0f61b70e.
//
// Solidity: function whiteListPool(address ) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) WhiteListPool(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListPool(&_OpvPool.CallOpts, arg0)
}

// WhiteListReceiver is a free data retrieval call binding the contract method 0x56e88ff1.
//
// Solidity: function whiteListReceiver(address ) view returns(bool)
func (_OpvPool *OpvPoolCaller) WhiteListReceiver(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "whiteListReceiver", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListReceiver is a free data retrieval call binding the contract method 0x56e88ff1.
//
// Solidity: function whiteListReceiver(address ) view returns(bool)
func (_OpvPool *OpvPoolSession) WhiteListReceiver(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListReceiver(&_OpvPool.CallOpts, arg0)
}

// WhiteListReceiver is a free data retrieval call binding the contract method 0x56e88ff1.
//
// Solidity: function whiteListReceiver(address ) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) WhiteListReceiver(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListReceiver(&_OpvPool.CallOpts, arg0)
}

// WhiteListSender is a free data retrieval call binding the contract method 0x64c0218f.
//
// Solidity: function whiteListSender(address ) view returns(bool)
func (_OpvPool *OpvPoolCaller) WhiteListSender(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvPool.contract.Call(opts, &out, "whiteListSender", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListSender is a free data retrieval call binding the contract method 0x64c0218f.
//
// Solidity: function whiteListSender(address ) view returns(bool)
func (_OpvPool *OpvPoolSession) WhiteListSender(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListSender(&_OpvPool.CallOpts, arg0)
}

// WhiteListSender is a free data retrieval call binding the contract method 0x64c0218f.
//
// Solidity: function whiteListSender(address ) view returns(bool)
func (_OpvPool *OpvPoolCallerSession) WhiteListSender(arg0 common.Address) (bool, error) {
	return _OpvPool.Contract.WhiteListSender(&_OpvPool.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Approve(&_OpvPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Approve(&_OpvPool.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_OpvPool *OpvPoolTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_OpvPool *OpvPoolSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Burn(&_OpvPool.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_OpvPool *OpvPoolTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Burn(&_OpvPool.TransactOpts, amount)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 feeTransfer) returns()
func (_OpvPool *OpvPoolTransactor) ChangeFee(opts *bind.TransactOpts, feeTransfer *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "changeFee", feeTransfer)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 feeTransfer) returns()
func (_OpvPool *OpvPoolSession) ChangeFee(feeTransfer *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.ChangeFee(&_OpvPool.TransactOpts, feeTransfer)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 feeTransfer) returns()
func (_OpvPool *OpvPoolTransactorSession) ChangeFee(feeTransfer *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.ChangeFee(&_OpvPool.TransactOpts, feeTransfer)
}

// ChangeFeeWallet is a paid mutator transaction binding the contract method 0x3e4d0310.
//
// Solidity: function changeFeeWallet(address feeWallet) returns()
func (_OpvPool *OpvPoolTransactor) ChangeFeeWallet(opts *bind.TransactOpts, feeWallet common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "changeFeeWallet", feeWallet)
}

// ChangeFeeWallet is a paid mutator transaction binding the contract method 0x3e4d0310.
//
// Solidity: function changeFeeWallet(address feeWallet) returns()
func (_OpvPool *OpvPoolSession) ChangeFeeWallet(feeWallet common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ChangeFeeWallet(&_OpvPool.TransactOpts, feeWallet)
}

// ChangeFeeWallet is a paid mutator transaction binding the contract method 0x3e4d0310.
//
// Solidity: function changeFeeWallet(address feeWallet) returns()
func (_OpvPool *OpvPoolTransactorSession) ChangeFeeWallet(feeWallet common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ChangeFeeWallet(&_OpvPool.TransactOpts, feeWallet)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OpvPool *OpvPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OpvPool *OpvPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.DecreaseAllowance(&_OpvPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.DecreaseAllowance(&_OpvPool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OpvPool *OpvPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OpvPool *OpvPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.IncreaseAllowance(&_OpvPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.IncreaseAllowance(&_OpvPool.TransactOpts, spender, addedValue)
}

// ModifyBlackList is a paid mutator transaction binding the contract method 0xd5e40649.
//
// Solidity: function modifyBlackList(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactor) ModifyBlackList(opts *bind.TransactOpts, newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "modifyBlackList", newAddr, removedAddr)
}

// ModifyBlackList is a paid mutator transaction binding the contract method 0xd5e40649.
//
// Solidity: function modifyBlackList(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolSession) ModifyBlackList(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyBlackList(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyBlackList is a paid mutator transaction binding the contract method 0xd5e40649.
//
// Solidity: function modifyBlackList(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactorSession) ModifyBlackList(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyBlackList(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListBot is a paid mutator transaction binding the contract method 0xcb045c61.
//
// Solidity: function modifyWhiteListBot(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactor) ModifyWhiteListBot(opts *bind.TransactOpts, newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "modifyWhiteListBot", newAddr, removedAddr)
}

// ModifyWhiteListBot is a paid mutator transaction binding the contract method 0xcb045c61.
//
// Solidity: function modifyWhiteListBot(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolSession) ModifyWhiteListBot(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListBot(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListBot is a paid mutator transaction binding the contract method 0xcb045c61.
//
// Solidity: function modifyWhiteListBot(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactorSession) ModifyWhiteListBot(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListBot(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListPool is a paid mutator transaction binding the contract method 0x7730e819.
//
// Solidity: function modifyWhiteListPool(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactor) ModifyWhiteListPool(opts *bind.TransactOpts, newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "modifyWhiteListPool", newAddr, removedAddr)
}

// ModifyWhiteListPool is a paid mutator transaction binding the contract method 0x7730e819.
//
// Solidity: function modifyWhiteListPool(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolSession) ModifyWhiteListPool(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListPool(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListPool is a paid mutator transaction binding the contract method 0x7730e819.
//
// Solidity: function modifyWhiteListPool(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactorSession) ModifyWhiteListPool(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListPool(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListReceiver is a paid mutator transaction binding the contract method 0x36807a53.
//
// Solidity: function modifyWhiteListReceiver(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactor) ModifyWhiteListReceiver(opts *bind.TransactOpts, newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "modifyWhiteListReceiver", newAddr, removedAddr)
}

// ModifyWhiteListReceiver is a paid mutator transaction binding the contract method 0x36807a53.
//
// Solidity: function modifyWhiteListReceiver(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolSession) ModifyWhiteListReceiver(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListReceiver(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListReceiver is a paid mutator transaction binding the contract method 0x36807a53.
//
// Solidity: function modifyWhiteListReceiver(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactorSession) ModifyWhiteListReceiver(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListReceiver(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListSender is a paid mutator transaction binding the contract method 0xc24df1f2.
//
// Solidity: function modifyWhiteListSender(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactor) ModifyWhiteListSender(opts *bind.TransactOpts, newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "modifyWhiteListSender", newAddr, removedAddr)
}

// ModifyWhiteListSender is a paid mutator transaction binding the contract method 0xc24df1f2.
//
// Solidity: function modifyWhiteListSender(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolSession) ModifyWhiteListSender(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListSender(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// ModifyWhiteListSender is a paid mutator transaction binding the contract method 0xc24df1f2.
//
// Solidity: function modifyWhiteListSender(address[] newAddr, address[] removedAddr) returns()
func (_OpvPool *OpvPoolTransactorSession) ModifyWhiteListSender(newAddr []common.Address, removedAddr []common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.ModifyWhiteListSender(&_OpvPool.TransactOpts, newAddr, removedAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvPool *OpvPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvPool *OpvPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpvPool.Contract.RenounceOwnership(&_OpvPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvPool *OpvPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpvPool.Contract.RenounceOwnership(&_OpvPool.TransactOpts)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xc2c7c03a.
//
// Solidity: function setAntiBot(bool _enable) returns()
func (_OpvPool *OpvPoolTransactor) SetAntiBot(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "setAntiBot", _enable)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xc2c7c03a.
//
// Solidity: function setAntiBot(bool _enable) returns()
func (_OpvPool *OpvPoolSession) SetAntiBot(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetAntiBot(&_OpvPool.TransactOpts, _enable)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xc2c7c03a.
//
// Solidity: function setAntiBot(bool _enable) returns()
func (_OpvPool *OpvPoolTransactorSession) SetAntiBot(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetAntiBot(&_OpvPool.TransactOpts, _enable)
}

// SetLimitSell is a paid mutator transaction binding the contract method 0x3875acfe.
//
// Solidity: function setLimitSell(bool _enable) returns()
func (_OpvPool *OpvPoolTransactor) SetLimitSell(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "setLimitSell", _enable)
}

// SetLimitSell is a paid mutator transaction binding the contract method 0x3875acfe.
//
// Solidity: function setLimitSell(bool _enable) returns()
func (_OpvPool *OpvPoolSession) SetLimitSell(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetLimitSell(&_OpvPool.TransactOpts, _enable)
}

// SetLimitSell is a paid mutator transaction binding the contract method 0x3875acfe.
//
// Solidity: function setLimitSell(bool _enable) returns()
func (_OpvPool *OpvPoolTransactorSession) SetLimitSell(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetLimitSell(&_OpvPool.TransactOpts, _enable)
}

// SetNumTokensSell is a paid mutator transaction binding the contract method 0x923a1c57.
//
// Solidity: function setNumTokensSell(uint256 numTokensSellToAddToLiquidity) returns(bool)
func (_OpvPool *OpvPoolTransactor) SetNumTokensSell(opts *bind.TransactOpts, numTokensSellToAddToLiquidity *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "setNumTokensSell", numTokensSellToAddToLiquidity)
}

// SetNumTokensSell is a paid mutator transaction binding the contract method 0x923a1c57.
//
// Solidity: function setNumTokensSell(uint256 numTokensSellToAddToLiquidity) returns(bool)
func (_OpvPool *OpvPoolSession) SetNumTokensSell(numTokensSellToAddToLiquidity *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.SetNumTokensSell(&_OpvPool.TransactOpts, numTokensSellToAddToLiquidity)
}

// SetNumTokensSell is a paid mutator transaction binding the contract method 0x923a1c57.
//
// Solidity: function setNumTokensSell(uint256 numTokensSellToAddToLiquidity) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) SetNumTokensSell(numTokensSellToAddToLiquidity *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.SetNumTokensSell(&_OpvPool.TransactOpts, numTokensSellToAddToLiquidity)
}

// SetSwapWhiteList is a paid mutator transaction binding the contract method 0x11bdfd14.
//
// Solidity: function setSwapWhiteList(bool _enable) returns()
func (_OpvPool *OpvPoolTransactor) SetSwapWhiteList(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "setSwapWhiteList", _enable)
}

// SetSwapWhiteList is a paid mutator transaction binding the contract method 0x11bdfd14.
//
// Solidity: function setSwapWhiteList(bool _enable) returns()
func (_OpvPool *OpvPoolSession) SetSwapWhiteList(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetSwapWhiteList(&_OpvPool.TransactOpts, _enable)
}

// SetSwapWhiteList is a paid mutator transaction binding the contract method 0x11bdfd14.
//
// Solidity: function setSwapWhiteList(bool _enable) returns()
func (_OpvPool *OpvPoolTransactorSession) SetSwapWhiteList(_enable bool) (*types.Transaction, error) {
	return _OpvPool.Contract.SetSwapWhiteList(&_OpvPool.TransactOpts, _enable)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Transfer(&_OpvPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.Transfer(&_OpvPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferFrom(&_OpvPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OpvPool *OpvPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferFrom(&_OpvPool.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvPool *OpvPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvPool *OpvPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferOwnership(&_OpvPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvPool *OpvPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferOwnership(&_OpvPool.TransactOpts, newOwner)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf640d508.
//
// Solidity: function transferToken(address coinAddress, uint256 value, address to) returns()
func (_OpvPool *OpvPoolTransactor) TransferToken(opts *bind.TransactOpts, coinAddress common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _OpvPool.contract.Transact(opts, "transferToken", coinAddress, value, to)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf640d508.
//
// Solidity: function transferToken(address coinAddress, uint256 value, address to) returns()
func (_OpvPool *OpvPoolSession) TransferToken(coinAddress common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferToken(&_OpvPool.TransactOpts, coinAddress, value, to)
}

// TransferToken is a paid mutator transaction binding the contract method 0xf640d508.
//
// Solidity: function transferToken(address coinAddress, uint256 value, address to) returns()
func (_OpvPool *OpvPoolTransactorSession) TransferToken(coinAddress common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _OpvPool.Contract.TransferToken(&_OpvPool.TransactOpts, coinAddress, value, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvPool *OpvPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvPool *OpvPoolSession) Receive() (*types.Transaction, error) {
	return _OpvPool.Contract.Receive(&_OpvPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OpvPool *OpvPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _OpvPool.Contract.Receive(&_OpvPool.TransactOpts)
}

// OpvPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OpvPool contract.
type OpvPoolApprovalIterator struct {
	Event *OpvPoolApproval // Event containing the contract specifics and raw log

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
func (it *OpvPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvPoolApproval)
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
		it.Event = new(OpvPoolApproval)
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
func (it *OpvPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvPoolApproval represents a Approval event raised by the OpvPool contract.
type OpvPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OpvPool *OpvPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OpvPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OpvPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OpvPoolApprovalIterator{contract: _OpvPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OpvPool *OpvPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OpvPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OpvPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvPoolApproval)
				if err := _OpvPool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OpvPool *OpvPoolFilterer) ParseApproval(log types.Log) (*OpvPoolApproval, error) {
	event := new(OpvPoolApproval)
	if err := _OpvPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OpvPool contract.
type OpvPoolOwnershipTransferredIterator struct {
	Event *OpvPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpvPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvPoolOwnershipTransferred)
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
		it.Event = new(OpvPoolOwnershipTransferred)
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
func (it *OpvPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvPoolOwnershipTransferred represents a OwnershipTransferred event raised by the OpvPool contract.
type OpvPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpvPool *OpvPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpvPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpvPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpvPoolOwnershipTransferredIterator{contract: _OpvPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpvPool *OpvPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpvPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpvPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvPoolOwnershipTransferred)
				if err := _OpvPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OpvPool *OpvPoolFilterer) ParseOwnershipTransferred(log types.Log) (*OpvPoolOwnershipTransferred, error) {
	event := new(OpvPoolOwnershipTransferred)
	if err := _OpvPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OpvPool contract.
type OpvPoolTransferIterator struct {
	Event *OpvPoolTransfer // Event containing the contract specifics and raw log

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
func (it *OpvPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvPoolTransfer)
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
		it.Event = new(OpvPoolTransfer)
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
func (it *OpvPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvPoolTransfer represents a Transfer event raised by the OpvPool contract.
type OpvPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OpvPool *OpvPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OpvPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpvPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OpvPoolTransferIterator{contract: _OpvPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OpvPool *OpvPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OpvPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpvPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvPoolTransfer)
				if err := _OpvPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OpvPool *OpvPoolFilterer) ParseTransfer(log types.Log) (*OpvPoolTransfer, error) {
	event := new(OpvPoolTransfer)
	if err := _OpvPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
