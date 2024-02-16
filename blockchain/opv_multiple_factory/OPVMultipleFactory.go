// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opv_multiple_factory

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

// OpvMultipleFactoryMetaData contains all meta data concerning the OpvMultipleFactory contract.
var OpvMultipleFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"defaultUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"marketplaceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opvAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uniqueId\",\"type\":\"uint256\"}],\"name\":\"LogBatchMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPV\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"addApprovalWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvalWhitelists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniqueId\",\"type\":\"uint256\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creatorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"removeApprovalWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"name\":\"updateBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OpvMultipleFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OpvMultipleFactoryMetaData.ABI instead.
var OpvMultipleFactoryABI = OpvMultipleFactoryMetaData.ABI

// OpvMultipleFactory is an auto generated Go binding around an Ethereum contract.
type OpvMultipleFactory struct {
	OpvMultipleFactoryCaller     // Read-only binding to the contract
	OpvMultipleFactoryTransactor // Write-only binding to the contract
	OpvMultipleFactoryFilterer   // Log filterer for contract events
}

// OpvMultipleFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpvMultipleFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvMultipleFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpvMultipleFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvMultipleFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpvMultipleFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpvMultipleFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpvMultipleFactorySession struct {
	Contract     *OpvMultipleFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OpvMultipleFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpvMultipleFactoryCallerSession struct {
	Contract *OpvMultipleFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OpvMultipleFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpvMultipleFactoryTransactorSession struct {
	Contract     *OpvMultipleFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OpvMultipleFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpvMultipleFactoryRaw struct {
	Contract *OpvMultipleFactory // Generic contract binding to access the raw methods on
}

// OpvMultipleFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpvMultipleFactoryCallerRaw struct {
	Contract *OpvMultipleFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OpvMultipleFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpvMultipleFactoryTransactorRaw struct {
	Contract *OpvMultipleFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpvMultipleFactory creates a new instance of OpvMultipleFactory, bound to a specific deployed contract.
func NewOpvMultipleFactory(address common.Address, backend bind.ContractBackend) (*OpvMultipleFactory, error) {
	contract, err := bindOpvMultipleFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactory{OpvMultipleFactoryCaller: OpvMultipleFactoryCaller{contract: contract}, OpvMultipleFactoryTransactor: OpvMultipleFactoryTransactor{contract: contract}, OpvMultipleFactoryFilterer: OpvMultipleFactoryFilterer{contract: contract}}, nil
}

// NewOpvMultipleFactoryCaller creates a new read-only instance of OpvMultipleFactory, bound to a specific deployed contract.
func NewOpvMultipleFactoryCaller(address common.Address, caller bind.ContractCaller) (*OpvMultipleFactoryCaller, error) {
	contract, err := bindOpvMultipleFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryCaller{contract: contract}, nil
}

// NewOpvMultipleFactoryTransactor creates a new write-only instance of OpvMultipleFactory, bound to a specific deployed contract.
func NewOpvMultipleFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OpvMultipleFactoryTransactor, error) {
	contract, err := bindOpvMultipleFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryTransactor{contract: contract}, nil
}

// NewOpvMultipleFactoryFilterer creates a new log filterer instance of OpvMultipleFactory, bound to a specific deployed contract.
func NewOpvMultipleFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OpvMultipleFactoryFilterer, error) {
	contract, err := bindOpvMultipleFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryFilterer{contract: contract}, nil
}

// bindOpvMultipleFactory binds a generic wrapper to an already deployed contract.
func bindOpvMultipleFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpvMultipleFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvMultipleFactory *OpvMultipleFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvMultipleFactory.Contract.OpvMultipleFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvMultipleFactory *OpvMultipleFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.OpvMultipleFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvMultipleFactory *OpvMultipleFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.OpvMultipleFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpvMultipleFactory *OpvMultipleFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpvMultipleFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpvMultipleFactory.Contract.DEFAULTADMINROLE(&_OpvMultipleFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpvMultipleFactory.Contract.DEFAULTADMINROLE(&_OpvMultipleFactory.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactorySession) MINTERROLE() ([32]byte, error) {
	return _OpvMultipleFactory.Contract.MINTERROLE(&_OpvMultipleFactory.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) MINTERROLE() ([32]byte, error) {
	return _OpvMultipleFactory.Contract.MINTERROLE(&_OpvMultipleFactory.CallOpts)
}

// OPV is a free data retrieval call binding the contract method 0x6906e075.
//
// Solidity: function OPV() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) OPV(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "OPV")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPV is a free data retrieval call binding the contract method 0x6906e075.
//
// Solidity: function OPV() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) OPV() (common.Address, error) {
	return _OpvMultipleFactory.Contract.OPV(&_OpvMultipleFactory.CallOpts)
}

// OPV is a free data retrieval call binding the contract method 0x6906e075.
//
// Solidity: function OPV() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) OPV() (common.Address, error) {
	return _OpvMultipleFactory.Contract.OPV(&_OpvMultipleFactory.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x0135f740.
//
// Solidity: function _feeAddress() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "_feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x0135f740.
//
// Solidity: function _feeAddress() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) FeeAddress() (common.Address, error) {
	return _OpvMultipleFactory.Contract.FeeAddress(&_OpvMultipleFactory.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x0135f740.
//
// Solidity: function _feeAddress() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) FeeAddress() (common.Address, error) {
	return _OpvMultipleFactory.Contract.FeeAddress(&_OpvMultipleFactory.CallOpts)
}

// FeeMint is a free data retrieval call binding the contract method 0xd942911f.
//
// Solidity: function _feeMint() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) FeeMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "_feeMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeMint is a free data retrieval call binding the contract method 0xd942911f.
//
// Solidity: function _feeMint() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) FeeMint() (*big.Int, error) {
	return _OpvMultipleFactory.Contract.FeeMint(&_OpvMultipleFactory.CallOpts)
}

// FeeMint is a free data retrieval call binding the contract method 0xd942911f.
//
// Solidity: function _feeMint() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) FeeMint() (*big.Int, error) {
	return _OpvMultipleFactory.Contract.FeeMint(&_OpvMultipleFactory.CallOpts)
}

// ApprovalWhitelists is a free data retrieval call binding the contract method 0x419e583d.
//
// Solidity: function approvalWhitelists(address ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) ApprovalWhitelists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "approvalWhitelists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalWhitelists is a free data retrieval call binding the contract method 0x419e583d.
//
// Solidity: function approvalWhitelists(address ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) ApprovalWhitelists(arg0 common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.ApprovalWhitelists(&_OpvMultipleFactory.CallOpts, arg0)
}

// ApprovalWhitelists is a free data retrieval call binding the contract method 0x419e583d.
//
// Solidity: function approvalWhitelists(address ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) ApprovalWhitelists(arg0 common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.ApprovalWhitelists(&_OpvMultipleFactory.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.BalanceOf(&_OpvMultipleFactory.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.BalanceOf(&_OpvMultipleFactory.CallOpts, owner)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) CreatorOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "creatorOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.CreatorOf(&_OpvMultipleFactory.CallOpts, tokenId)
}

// CreatorOf is a free data retrieval call binding the contract method 0x589a1743.
//
// Solidity: function creatorOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) CreatorOf(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.CreatorOf(&_OpvMultipleFactory.CallOpts, tokenId)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) Creators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "creators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) Creators(arg0 *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.Creators(&_OpvMultipleFactory.CallOpts, arg0)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.Creators(&_OpvMultipleFactory.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.GetApproved(&_OpvMultipleFactory.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.GetApproved(&_OpvMultipleFactory.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpvMultipleFactory.Contract.GetRoleAdmin(&_OpvMultipleFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpvMultipleFactory.Contract.GetRoleAdmin(&_OpvMultipleFactory.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.GetRoleMember(&_OpvMultipleFactory.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.GetRoleMember(&_OpvMultipleFactory.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.GetRoleMemberCount(&_OpvMultipleFactory.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.GetRoleMemberCount(&_OpvMultipleFactory.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.HasRole(&_OpvMultipleFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.HasRole(&_OpvMultipleFactory.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.IsApprovedForAll(&_OpvMultipleFactory.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _OpvMultipleFactory.Contract.IsApprovedForAll(&_OpvMultipleFactory.CallOpts, owner, operator)
}

// IsLocked is a free data retrieval call binding the contract method 0xf6aacfb1.
//
// Solidity: function isLocked(uint256 tokenId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) IsLocked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "isLocked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xf6aacfb1.
//
// Solidity: function isLocked(uint256 tokenId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) IsLocked(tokenId *big.Int) (bool, error) {
	return _OpvMultipleFactory.Contract.IsLocked(&_OpvMultipleFactory.CallOpts, tokenId)
}

// IsLocked is a free data retrieval call binding the contract method 0xf6aacfb1.
//
// Solidity: function isLocked(uint256 tokenId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) IsLocked(tokenId *big.Int) (bool, error) {
	return _OpvMultipleFactory.Contract.IsLocked(&_OpvMultipleFactory.CallOpts, tokenId)
}

// LockedTokens is a free data retrieval call binding the contract method 0xdcec3294.
//
// Solidity: function lockedTokens(uint256 ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) LockedTokens(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "lockedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockedTokens is a free data retrieval call binding the contract method 0xdcec3294.
//
// Solidity: function lockedTokens(uint256 ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) LockedTokens(arg0 *big.Int) (bool, error) {
	return _OpvMultipleFactory.Contract.LockedTokens(&_OpvMultipleFactory.CallOpts, arg0)
}

// LockedTokens is a free data retrieval call binding the contract method 0xdcec3294.
//
// Solidity: function lockedTokens(uint256 ) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) LockedTokens(arg0 *big.Int) (bool, error) {
	return _OpvMultipleFactory.Contract.LockedTokens(&_OpvMultipleFactory.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactorySession) Name() (string, error) {
	return _OpvMultipleFactory.Contract.Name(&_OpvMultipleFactory.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) Name() (string, error) {
	return _OpvMultipleFactory.Contract.Name(&_OpvMultipleFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) Owner() (common.Address, error) {
	return _OpvMultipleFactory.Contract.Owner(&_OpvMultipleFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) Owner() (common.Address, error) {
	return _OpvMultipleFactory.Contract.Owner(&_OpvMultipleFactory.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactorySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.OwnerOf(&_OpvMultipleFactory.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _OpvMultipleFactory.Contract.OwnerOf(&_OpvMultipleFactory.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpvMultipleFactory.Contract.SupportsInterface(&_OpvMultipleFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpvMultipleFactory.Contract.SupportsInterface(&_OpvMultipleFactory.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactorySession) Symbol() (string, error) {
	return _OpvMultipleFactory.Contract.Symbol(&_OpvMultipleFactory.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) Symbol() (string, error) {
	return _OpvMultipleFactory.Contract.Symbol(&_OpvMultipleFactory.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TokenByIndex(&_OpvMultipleFactory.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TokenByIndex(&_OpvMultipleFactory.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TokenOfOwnerByIndex(&_OpvMultipleFactory.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TokenOfOwnerByIndex(&_OpvMultipleFactory.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactorySession) TokenURI(tokenId *big.Int) (string, error) {
	return _OpvMultipleFactory.Contract.TokenURI(&_OpvMultipleFactory.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _OpvMultipleFactory.Contract.TokenURI(&_OpvMultipleFactory.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpvMultipleFactory.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactorySession) TotalSupply() (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TotalSupply(&_OpvMultipleFactory.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OpvMultipleFactory *OpvMultipleFactoryCallerSession) TotalSupply() (*big.Int, error) {
	return _OpvMultipleFactory.Contract.TotalSupply(&_OpvMultipleFactory.CallOpts)
}

// AddApprovalWhitelist is a paid mutator transaction binding the contract method 0xb78437ea.
//
// Solidity: function addApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) AddApprovalWhitelist(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "addApprovalWhitelist", proxy)
}

// AddApprovalWhitelist is a paid mutator transaction binding the contract method 0xb78437ea.
//
// Solidity: function addApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) AddApprovalWhitelist(proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.AddApprovalWhitelist(&_OpvMultipleFactory.TransactOpts, proxy)
}

// AddApprovalWhitelist is a paid mutator transaction binding the contract method 0xb78437ea.
//
// Solidity: function addApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) AddApprovalWhitelist(proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.AddApprovalWhitelist(&_OpvMultipleFactory.TransactOpts, proxy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Approve(&_OpvMultipleFactory.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Approve(&_OpvMultipleFactory.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x2a959b89.
//
// Solidity: function batchMint(address to, uint256 amount, uint256 uniqueId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) BatchMint(opts *bind.TransactOpts, to common.Address, amount *big.Int, uniqueId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "batchMint", to, amount, uniqueId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x2a959b89.
//
// Solidity: function batchMint(address to, uint256 amount, uint256 uniqueId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) BatchMint(to common.Address, amount *big.Int, uniqueId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.BatchMint(&_OpvMultipleFactory.TransactOpts, to, amount, uniqueId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x2a959b89.
//
// Solidity: function batchMint(address to, uint256 amount, uint256 uniqueId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) BatchMint(to common.Address, amount *big.Int, uniqueId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.BatchMint(&_OpvMultipleFactory.TransactOpts, to, amount, uniqueId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.GrantRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.GrantRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) Lock(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "lock", tokenId)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) Lock(tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Lock(&_OpvMultipleFactory.TransactOpts, tokenId)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) Lock(tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Lock(&_OpvMultipleFactory.TransactOpts, tokenId)
}

// RemoveApprovalWhitelist is a paid mutator transaction binding the contract method 0x0a75eea2.
//
// Solidity: function removeApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) RemoveApprovalWhitelist(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "removeApprovalWhitelist", proxy)
}

// RemoveApprovalWhitelist is a paid mutator transaction binding the contract method 0x0a75eea2.
//
// Solidity: function removeApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) RemoveApprovalWhitelist(proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RemoveApprovalWhitelist(&_OpvMultipleFactory.TransactOpts, proxy)
}

// RemoveApprovalWhitelist is a paid mutator transaction binding the contract method 0x0a75eea2.
//
// Solidity: function removeApprovalWhitelist(address proxy) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) RemoveApprovalWhitelist(proxy common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RemoveApprovalWhitelist(&_OpvMultipleFactory.TransactOpts, proxy)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RenounceOwnership(&_OpvMultipleFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RenounceOwnership(&_OpvMultipleFactory.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RenounceRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RenounceRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RevokeRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.RevokeRole(&_OpvMultipleFactory.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SafeTransferFrom(&_OpvMultipleFactory.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SafeTransferFrom(&_OpvMultipleFactory.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SafeTransferFrom0(&_OpvMultipleFactory.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SafeTransferFrom0(&_OpvMultipleFactory.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetApprovalForAll(&_OpvMultipleFactory.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetApprovalForAll(&_OpvMultipleFactory.TransactOpts, operator, approved)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 amount) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) SetFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "setFee", amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 amount) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) SetFee(amount *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetFee(&_OpvMultipleFactory.TransactOpts, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 amount) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) SetFee(amount *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetFee(&_OpvMultipleFactory.TransactOpts, amount)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address feeAddress) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) SetFeeAddress(opts *bind.TransactOpts, feeAddress common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "setFeeAddress", feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address feeAddress) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) SetFeeAddress(feeAddress common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetFeeAddress(&_OpvMultipleFactory.TransactOpts, feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address feeAddress) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) SetFeeAddress(feeAddress common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.SetFeeAddress(&_OpvMultipleFactory.TransactOpts, feeAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.TransferFrom(&_OpvMultipleFactory.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.TransferFrom(&_OpvMultipleFactory.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.TransferOwnership(&_OpvMultipleFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.TransferOwnership(&_OpvMultipleFactory.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) Unlock(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "unlock", tokenId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) Unlock(tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Unlock(&_OpvMultipleFactory.TransactOpts, tokenId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 tokenId) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) Unlock(tokenId *big.Int) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.Unlock(&_OpvMultipleFactory.TransactOpts, tokenId)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string baseTokenURI) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactor) UpdateBaseURI(opts *bind.TransactOpts, baseTokenURI string) (*types.Transaction, error) {
	return _OpvMultipleFactory.contract.Transact(opts, "updateBaseURI", baseTokenURI)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string baseTokenURI) returns()
func (_OpvMultipleFactory *OpvMultipleFactorySession) UpdateBaseURI(baseTokenURI string) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.UpdateBaseURI(&_OpvMultipleFactory.TransactOpts, baseTokenURI)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string baseTokenURI) returns()
func (_OpvMultipleFactory *OpvMultipleFactoryTransactorSession) UpdateBaseURI(baseTokenURI string) (*types.Transaction, error) {
	return _OpvMultipleFactory.Contract.UpdateBaseURI(&_OpvMultipleFactory.TransactOpts, baseTokenURI)
}

// OpvMultipleFactoryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryApprovalIterator struct {
	Event *OpvMultipleFactoryApproval // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryApproval)
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
		it.Event = new(OpvMultipleFactoryApproval)
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
func (it *OpvMultipleFactoryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryApproval represents a Approval event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*OpvMultipleFactoryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryApprovalIterator{contract: _OpvMultipleFactory.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryApproval)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseApproval(log types.Log) (*OpvMultipleFactoryApproval, error) {
	event := new(OpvMultipleFactoryApproval)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryApprovalForAllIterator struct {
	Event *OpvMultipleFactoryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryApprovalForAll)
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
		it.Event = new(OpvMultipleFactoryApprovalForAll)
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
func (it *OpvMultipleFactoryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryApprovalForAll represents a ApprovalForAll event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*OpvMultipleFactoryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryApprovalForAllIterator{contract: _OpvMultipleFactory.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryApprovalForAll)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseApprovalForAll(log types.Log) (*OpvMultipleFactoryApprovalForAll, error) {
	event := new(OpvMultipleFactoryApprovalForAll)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryLogBatchMintIterator is returned from FilterLogBatchMint and is used to iterate over the raw logs and unpacked data for LogBatchMint events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryLogBatchMintIterator struct {
	Event *OpvMultipleFactoryLogBatchMint // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryLogBatchMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryLogBatchMint)
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
		it.Event = new(OpvMultipleFactoryLogBatchMint)
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
func (it *OpvMultipleFactoryLogBatchMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryLogBatchMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryLogBatchMint represents a LogBatchMint event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryLogBatchMint struct {
	User     common.Address
	FromId   *big.Int
	ToId     *big.Int
	DataId   *big.Int
	UniqueId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogBatchMint is a free log retrieval operation binding the contract event 0xe7c9f23f1a2a2c3e7d20c5376d1bb99d93ca64e37c189f70dab2a3d0289dec69.
//
// Solidity: event LogBatchMint(address user, uint256 fromId, uint256 toId, uint256 dataId, uint256 uniqueId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterLogBatchMint(opts *bind.FilterOpts) (*OpvMultipleFactoryLogBatchMintIterator, error) {

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "LogBatchMint")
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryLogBatchMintIterator{contract: _OpvMultipleFactory.contract, event: "LogBatchMint", logs: logs, sub: sub}, nil
}

// WatchLogBatchMint is a free log subscription operation binding the contract event 0xe7c9f23f1a2a2c3e7d20c5376d1bb99d93ca64e37c189f70dab2a3d0289dec69.
//
// Solidity: event LogBatchMint(address user, uint256 fromId, uint256 toId, uint256 dataId, uint256 uniqueId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchLogBatchMint(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryLogBatchMint) (event.Subscription, error) {

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "LogBatchMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryLogBatchMint)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "LogBatchMint", log); err != nil {
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

// ParseLogBatchMint is a log parse operation binding the contract event 0xe7c9f23f1a2a2c3e7d20c5376d1bb99d93ca64e37c189f70dab2a3d0289dec69.
//
// Solidity: event LogBatchMint(address user, uint256 fromId, uint256 toId, uint256 dataId, uint256 uniqueId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseLogBatchMint(log types.Log) (*OpvMultipleFactoryLogBatchMint, error) {
	event := new(OpvMultipleFactoryLogBatchMint)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "LogBatchMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryOwnershipTransferredIterator struct {
	Event *OpvMultipleFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryOwnershipTransferred)
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
		it.Event = new(OpvMultipleFactoryOwnershipTransferred)
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
func (it *OpvMultipleFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpvMultipleFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryOwnershipTransferredIterator{contract: _OpvMultipleFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryOwnershipTransferred)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*OpvMultipleFactoryOwnershipTransferred, error) {
	event := new(OpvMultipleFactoryOwnershipTransferred)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleAdminChangedIterator struct {
	Event *OpvMultipleFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryRoleAdminChanged)
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
		it.Event = new(OpvMultipleFactoryRoleAdminChanged)
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
func (it *OpvMultipleFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OpvMultipleFactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryRoleAdminChangedIterator{contract: _OpvMultipleFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryRoleAdminChanged)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*OpvMultipleFactoryRoleAdminChanged, error) {
	event := new(OpvMultipleFactoryRoleAdminChanged)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleGrantedIterator struct {
	Event *OpvMultipleFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryRoleGranted)
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
		it.Event = new(OpvMultipleFactoryRoleGranted)
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
func (it *OpvMultipleFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryRoleGranted represents a RoleGranted event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpvMultipleFactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryRoleGrantedIterator{contract: _OpvMultipleFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryRoleGranted)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseRoleGranted(log types.Log) (*OpvMultipleFactoryRoleGranted, error) {
	event := new(OpvMultipleFactoryRoleGranted)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleRevokedIterator struct {
	Event *OpvMultipleFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryRoleRevoked)
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
		it.Event = new(OpvMultipleFactoryRoleRevoked)
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
func (it *OpvMultipleFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryRoleRevoked represents a RoleRevoked event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpvMultipleFactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryRoleRevokedIterator{contract: _OpvMultipleFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryRoleRevoked)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseRoleRevoked(log types.Log) (*OpvMultipleFactoryRoleRevoked, error) {
	event := new(OpvMultipleFactoryRoleRevoked)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpvMultipleFactoryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryTransferIterator struct {
	Event *OpvMultipleFactoryTransfer // Event containing the contract specifics and raw log

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
func (it *OpvMultipleFactoryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpvMultipleFactoryTransfer)
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
		it.Event = new(OpvMultipleFactoryTransfer)
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
func (it *OpvMultipleFactoryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpvMultipleFactoryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpvMultipleFactoryTransfer represents a Transfer event raised by the OpvMultipleFactory contract.
type OpvMultipleFactoryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*OpvMultipleFactoryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OpvMultipleFactoryTransferIterator{contract: _OpvMultipleFactory.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OpvMultipleFactoryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _OpvMultipleFactory.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpvMultipleFactoryTransfer)
				if err := _OpvMultipleFactory.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_OpvMultipleFactory *OpvMultipleFactoryFilterer) ParseTransfer(log types.Log) (*OpvMultipleFactoryTransfer, error) {
	event := new(OpvMultipleFactoryTransfer)
	if err := _OpvMultipleFactory.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
