// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_version\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetTokenInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Minter is a free data retrieval call binding the contract method 0x578ec33f.
//
// Solidity: function _minter() view returns(address)
func (_Storage *StorageCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "_minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x578ec33f.
//
// Solidity: function _minter() view returns(address)
func (_Storage *StorageSession) Minter() (common.Address, error) {
	return _Storage.Contract.Minter(&_Storage.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x578ec33f.
//
// Solidity: function _minter() view returns(address)
func (_Storage *StorageCallerSession) Minter() (common.Address, error) {
	return _Storage.Contract.Minter(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x3e118dbe.
//
// Solidity: function _version() view returns(string name, string version)
func (_Storage *StorageCaller) Version(opts *bind.CallOpts) (struct {
	Name    string
	Version string
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "_version")

	outstruct := new(struct {
		Name    string
		Version string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Version is a free data retrieval call binding the contract method 0x3e118dbe.
//
// Solidity: function _version() view returns(string name, string version)
func (_Storage *StorageSession) Version() (struct {
	Name    string
	Version string
}, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x3e118dbe.
//
// Solidity: function _version() view returns(string name, string version)
func (_Storage *StorageCallerSession) Version() (struct {
	Name    string
	Version string
}, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Storage *StorageCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Storage *StorageCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCallerSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Storage *StorageCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Storage *StorageSession) IsInitialized() (bool, error) {
	return _Storage.Contract.IsInitialized(&_Storage.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Storage *StorageCallerSession) IsInitialized() (bool, error) {
	return _Storage.Contract.IsInitialized(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCallerSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Storage.Contract.SupportsInterface(&_Storage.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Storage *StorageCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Storage.Contract.SupportsInterface(&_Storage.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCallerSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCallerSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Storage *StorageSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, spender, value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "decreaseApproval", spender, value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageSession) DecreaseApproval(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecreaseApproval(&_Storage.TransactOpts, spender, value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactorSession) DecreaseApproval(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecreaseApproval(&_Storage.TransactOpts, spender, value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "increaseApproval", spender, value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageSession) IncreaseApproval(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.IncreaseApproval(&_Storage.TransactOpts, spender, value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 value) returns(bool)
func (_Storage *StorageTransactorSession) IncreaseApproval(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.IncreaseApproval(&_Storage.TransactOpts, spender, value)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Storage *StorageTransactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Storage *StorageSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransfer(&_Storage.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Storage *StorageTransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransfer(&_Storage.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageTransactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransfer0(&_Storage.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageTransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransfer0(&_Storage.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Storage *StorageTransactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Storage *StorageSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransferFrom(&_Storage.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Storage *StorageTransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransferFrom(&_Storage.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageTransactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransferFrom0(&_Storage.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Storage *StorageTransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SafeTransferFrom0(&_Storage.TransactOpts, sender, recipient, amount, data)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address minter) returns()
func (_Storage *StorageTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setMinter", minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address minter) returns()
func (_Storage *StorageSession) SetMinter(minter common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetMinter(&_Storage.TransactOpts, minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address minter) returns()
func (_Storage *StorageTransactorSession) SetMinter(minter common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetMinter(&_Storage.TransactOpts, minter)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Storage *StorageTransactor) SetOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setOwner", owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Storage *StorageSession) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetOwner(&_Storage.TransactOpts, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Storage *StorageTransactorSession) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetOwner(&_Storage.TransactOpts, owner)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0xda262f58.
//
// Solidity: function setTokenInfo(string tokenName, string tokenSymbol) returns()
func (_Storage *StorageTransactor) SetTokenInfo(opts *bind.TransactOpts, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setTokenInfo", tokenName, tokenSymbol)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0xda262f58.
//
// Solidity: function setTokenInfo(string tokenName, string tokenSymbol) returns()
func (_Storage *StorageSession) SetTokenInfo(tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _Storage.Contract.SetTokenInfo(&_Storage.TransactOpts, tokenName, tokenSymbol)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0xda262f58.
//
// Solidity: function setTokenInfo(string tokenName, string tokenSymbol) returns()
func (_Storage *StorageTransactorSession) SetTokenInfo(tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _Storage.Contract.SetTokenInfo(&_Storage.TransactOpts, tokenName, tokenSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Storage *StorageSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Storage *StorageTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Storage *StorageTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Storage *StorageSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Storage *StorageTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, sender, recipient, amount)
}

// StorageApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Storage contract.
type StorageApprovalIterator struct {
	Event *StorageApproval // Event containing the contract specifics and raw log

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
func (it *StorageApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageApproval)
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
		it.Event = new(StorageApproval)
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
func (it *StorageApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageApproval represents a Approval event raised by the Storage contract.
type StorageApproval struct {
	Holder  common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed holder, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) FilterApproval(opts *bind.FilterOpts, holder []common.Address, spender []common.Address) (*StorageApprovalIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Approval", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StorageApprovalIterator{contract: _Storage.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed holder, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StorageApproval, holder []common.Address, spender []common.Address) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Approval", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageApproval)
				if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed holder, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) ParseApproval(log types.Log) (*StorageApproval, error) {
	event := new(StorageApproval)
	if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the Storage contract.
type StorageSetMinterIterator struct {
	Event *StorageSetMinter // Event containing the contract specifics and raw log

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
func (it *StorageSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetMinter)
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
		it.Event = new(StorageSetMinter)
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
func (it *StorageSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetMinter represents a SetMinter event raised by the Storage contract.
type StorageSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Storage *StorageFilterer) FilterSetMinter(opts *bind.FilterOpts) (*StorageSetMinterIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return &StorageSetMinterIterator{contract: _Storage.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Storage *StorageFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *StorageSetMinter) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetMinter)
				if err := _Storage.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Storage *StorageFilterer) ParseSetMinter(log types.Log) (*StorageSetMinter, error) {
	event := new(StorageSetMinter)
	if err := _Storage.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the Storage contract.
type StorageSetOwnerIterator struct {
	Event *StorageSetOwner // Event containing the contract specifics and raw log

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
func (it *StorageSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetOwner)
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
		it.Event = new(StorageSetOwner)
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
func (it *StorageSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetOwner represents a SetOwner event raised by the Storage contract.
type StorageSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address owner)
func (_Storage *StorageFilterer) FilterSetOwner(opts *bind.FilterOpts) (*StorageSetOwnerIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetOwner")
	if err != nil {
		return nil, err
	}
	return &StorageSetOwnerIterator{contract: _Storage.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address owner)
func (_Storage *StorageFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *StorageSetOwner) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetOwner)
				if err := _Storage.contract.UnpackLog(event, "SetOwner", log); err != nil {
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

// ParseSetOwner is a log parse operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address owner)
func (_Storage *StorageFilterer) ParseSetOwner(log types.Log) (*StorageSetOwner, error) {
	event := new(StorageSetOwner)
	if err := _Storage.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetTokenInfoIterator is returned from FilterSetTokenInfo and is used to iterate over the raw logs and unpacked data for SetTokenInfo events raised by the Storage contract.
type StorageSetTokenInfoIterator struct {
	Event *StorageSetTokenInfo // Event containing the contract specifics and raw log

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
func (it *StorageSetTokenInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetTokenInfo)
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
		it.Event = new(StorageSetTokenInfo)
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
func (it *StorageSetTokenInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetTokenInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetTokenInfo represents a SetTokenInfo event raised by the Storage contract.
type StorageSetTokenInfo struct {
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTokenInfo is a free log retrieval operation binding the contract event 0x9a6645cd07452c704655abb8d3df1bb29cc145ac1bdd0cb6370469cf6bbc9a29.
//
// Solidity: event SetTokenInfo(string name, string symbol)
func (_Storage *StorageFilterer) FilterSetTokenInfo(opts *bind.FilterOpts) (*StorageSetTokenInfoIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetTokenInfo")
	if err != nil {
		return nil, err
	}
	return &StorageSetTokenInfoIterator{contract: _Storage.contract, event: "SetTokenInfo", logs: logs, sub: sub}, nil
}

// WatchSetTokenInfo is a free log subscription operation binding the contract event 0x9a6645cd07452c704655abb8d3df1bb29cc145ac1bdd0cb6370469cf6bbc9a29.
//
// Solidity: event SetTokenInfo(string name, string symbol)
func (_Storage *StorageFilterer) WatchSetTokenInfo(opts *bind.WatchOpts, sink chan<- *StorageSetTokenInfo) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetTokenInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetTokenInfo)
				if err := _Storage.contract.UnpackLog(event, "SetTokenInfo", log); err != nil {
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

// ParseSetTokenInfo is a log parse operation binding the contract event 0x9a6645cd07452c704655abb8d3df1bb29cc145ac1bdd0cb6370469cf6bbc9a29.
//
// Solidity: event SetTokenInfo(string name, string symbol)
func (_Storage *StorageFilterer) ParseSetTokenInfo(log types.Log) (*StorageSetTokenInfo, error) {
	event := new(StorageSetTokenInfo)
	if err := _Storage.contract.UnpackLog(event, "SetTokenInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Storage contract.
type StorageTransferIterator struct {
	Event *StorageTransfer // Event containing the contract specifics and raw log

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
func (it *StorageTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageTransfer)
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
		it.Event = new(StorageTransfer)
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
func (it *StorageTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageTransfer represents a Transfer event raised by the Storage contract.
type StorageTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StorageTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StorageTransferIterator{contract: _Storage.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StorageTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageTransfer)
				if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) ParseTransfer(log types.Log) (*StorageTransfer, error) {
	event := new(StorageTransfer)
	if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
