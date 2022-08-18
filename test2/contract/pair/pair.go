// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"addKlayLiquidityWithLimit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"grabKlayFromFactory\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addKlayLiquidity\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mining\",\"type\":\"uint256\"}],\"name\":\"changeMiningRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeNeg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mining\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amountA\",\"type\":\"uint256\"},{\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"addKctLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastMined\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"estimateNeg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateMiningIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangePos\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amountA\",\"type\":\"uint256\"},{\"name\":\"amountB\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"addKctLiquidityWithLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"estimatePos\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLastIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"initPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityWithLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_mining\",\"type\":\"uint256\"}],\"name\":\"ChangeMiningRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lastMined\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"miningIndex\",\"type\":\"uint256\"}],\"name\":\"UpdateMiningIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardSum\",\"type\":\"uint256\"}],\"name\":\"GiveReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"ExchangePos\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"ExchangeNeg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Storage *StorageCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Storage *StorageSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Storage *StorageCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, arg0)
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

// EstimateNeg is a free data retrieval call binding the contract method 0xc0759842.
//
// Solidity: function estimateNeg(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageCaller) EstimateNeg(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "estimateNeg", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateNeg is a free data retrieval call binding the contract method 0xc0759842.
//
// Solidity: function estimateNeg(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageSession) EstimateNeg(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.EstimateNeg(&_Storage.CallOpts, token, amount)
}

// EstimateNeg is a free data retrieval call binding the contract method 0xc0759842.
//
// Solidity: function estimateNeg(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageCallerSession) EstimateNeg(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.EstimateNeg(&_Storage.CallOpts, token, amount)
}

// EstimatePos is a free data retrieval call binding the contract method 0xe4161181.
//
// Solidity: function estimatePos(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageCaller) EstimatePos(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "estimatePos", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatePos is a free data retrieval call binding the contract method 0xe4161181.
//
// Solidity: function estimatePos(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageSession) EstimatePos(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.EstimatePos(&_Storage.CallOpts, token, amount)
}

// EstimatePos is a free data retrieval call binding the contract method 0xe4161181.
//
// Solidity: function estimatePos(address token, uint256 amount) view returns(uint256)
func (_Storage *StorageCallerSession) EstimatePos(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.EstimatePos(&_Storage.CallOpts, token, amount)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageSession) Factory() (common.Address, error) {
	return _Storage.Contract.Factory(&_Storage.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageCallerSession) Factory() (common.Address, error) {
	return _Storage.Contract.Factory(&_Storage.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Storage *StorageCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Storage *StorageSession) Fee() (*big.Int, error) {
	return _Storage.Contract.Fee(&_Storage.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Storage *StorageCallerSession) Fee() (*big.Int, error) {
	return _Storage.Contract.Fee(&_Storage.CallOpts)
}

// GetCurrentPool is a free data retrieval call binding the contract method 0x1a595f65.
//
// Solidity: function getCurrentPool() view returns(uint256, uint256)
func (_Storage *StorageCaller) GetCurrentPool(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCurrentPool")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCurrentPool is a free data retrieval call binding the contract method 0x1a595f65.
//
// Solidity: function getCurrentPool() view returns(uint256, uint256)
func (_Storage *StorageSession) GetCurrentPool() (*big.Int, *big.Int, error) {
	return _Storage.Contract.GetCurrentPool(&_Storage.CallOpts)
}

// GetCurrentPool is a free data retrieval call binding the contract method 0x1a595f65.
//
// Solidity: function getCurrentPool() view returns(uint256, uint256)
func (_Storage *StorageCallerSession) GetCurrentPool() (*big.Int, *big.Int, error) {
	return _Storage.Contract.GetCurrentPool(&_Storage.CallOpts)
}

// LastMined is a free data retrieval call binding the contract method 0xb5afac5d.
//
// Solidity: function lastMined() view returns(uint256)
func (_Storage *StorageCaller) LastMined(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "lastMined")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMined is a free data retrieval call binding the contract method 0xb5afac5d.
//
// Solidity: function lastMined() view returns(uint256)
func (_Storage *StorageSession) LastMined() (*big.Int, error) {
	return _Storage.Contract.LastMined(&_Storage.CallOpts)
}

// LastMined is a free data retrieval call binding the contract method 0xb5afac5d.
//
// Solidity: function lastMined() view returns(uint256)
func (_Storage *StorageCallerSession) LastMined() (*big.Int, error) {
	return _Storage.Contract.LastMined(&_Storage.CallOpts)
}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(uint256)
func (_Storage *StorageCaller) Mining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "mining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(uint256)
func (_Storage *StorageSession) Mining() (*big.Int, error) {
	return _Storage.Contract.Mining(&_Storage.CallOpts)
}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(uint256)
func (_Storage *StorageCallerSession) Mining() (*big.Int, error) {
	return _Storage.Contract.Mining(&_Storage.CallOpts)
}

// MiningIndex is a free data retrieval call binding the contract method 0x8d80fc0c.
//
// Solidity: function miningIndex() view returns(uint256)
func (_Storage *StorageCaller) MiningIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "miningIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningIndex is a free data retrieval call binding the contract method 0x8d80fc0c.
//
// Solidity: function miningIndex() view returns(uint256)
func (_Storage *StorageSession) MiningIndex() (*big.Int, error) {
	return _Storage.Contract.MiningIndex(&_Storage.CallOpts)
}

// MiningIndex is a free data retrieval call binding the contract method 0x8d80fc0c.
//
// Solidity: function miningIndex() view returns(uint256)
func (_Storage *StorageCallerSession) MiningIndex() (*big.Int, error) {
	return _Storage.Contract.MiningIndex(&_Storage.CallOpts)
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

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Storage *StorageCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Storage *StorageSession) TokenA() (common.Address, error) {
	return _Storage.Contract.TokenA(&_Storage.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Storage *StorageCallerSession) TokenA() (common.Address, error) {
	return _Storage.Contract.TokenA(&_Storage.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Storage *StorageCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Storage *StorageSession) TokenB() (common.Address, error) {
	return _Storage.Contract.TokenB(&_Storage.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Storage *StorageCallerSession) TokenB() (common.Address, error) {
	return _Storage.Contract.TokenB(&_Storage.CallOpts)
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

// UserLastIndex is a free data retrieval call binding the contract method 0xe7180f9a.
//
// Solidity: function userLastIndex(address ) view returns(uint256)
func (_Storage *StorageCaller) UserLastIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "userLastIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastIndex is a free data retrieval call binding the contract method 0xe7180f9a.
//
// Solidity: function userLastIndex(address ) view returns(uint256)
func (_Storage *StorageSession) UserLastIndex(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.UserLastIndex(&_Storage.CallOpts, arg0)
}

// UserLastIndex is a free data retrieval call binding the contract method 0xe7180f9a.
//
// Solidity: function userLastIndex(address ) view returns(uint256)
func (_Storage *StorageCallerSession) UserLastIndex(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.UserLastIndex(&_Storage.CallOpts, arg0)
}

// UserRewardSum is a free data retrieval call binding the contract method 0x5c307085.
//
// Solidity: function userRewardSum(address ) view returns(uint256)
func (_Storage *StorageCaller) UserRewardSum(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "userRewardSum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardSum is a free data retrieval call binding the contract method 0x5c307085.
//
// Solidity: function userRewardSum(address ) view returns(uint256)
func (_Storage *StorageSession) UserRewardSum(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.UserRewardSum(&_Storage.CallOpts, arg0)
}

// UserRewardSum is a free data retrieval call binding the contract method 0x5c307085.
//
// Solidity: function userRewardSum(address ) view returns(uint256)
func (_Storage *StorageCallerSession) UserRewardSum(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.UserRewardSum(&_Storage.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageSession) Version() (string, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageCallerSession) Version() (string, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// AddKctLiquidity is a paid mutator transaction binding the contract method 0xabeaa7b6.
//
// Solidity: function addKctLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Storage *StorageTransactor) AddKctLiquidity(opts *bind.TransactOpts, amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addKctLiquidity", amountA, amountB)
}

// AddKctLiquidity is a paid mutator transaction binding the contract method 0xabeaa7b6.
//
// Solidity: function addKctLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Storage *StorageSession) AddKctLiquidity(amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKctLiquidity(&_Storage.TransactOpts, amountA, amountB)
}

// AddKctLiquidity is a paid mutator transaction binding the contract method 0xabeaa7b6.
//
// Solidity: function addKctLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Storage *StorageTransactorSession) AddKctLiquidity(amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKctLiquidity(&_Storage.TransactOpts, amountA, amountB)
}

// AddKctLiquidityWithLimit is a paid mutator transaction binding the contract method 0xd9f25e88.
//
// Solidity: function addKctLiquidityWithLimit(uint256 amountA, uint256 amountB, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageTransactor) AddKctLiquidityWithLimit(opts *bind.TransactOpts, amountA *big.Int, amountB *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addKctLiquidityWithLimit", amountA, amountB, minAmountA, minAmountB)
}

// AddKctLiquidityWithLimit is a paid mutator transaction binding the contract method 0xd9f25e88.
//
// Solidity: function addKctLiquidityWithLimit(uint256 amountA, uint256 amountB, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageSession) AddKctLiquidityWithLimit(amountA *big.Int, amountB *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKctLiquidityWithLimit(&_Storage.TransactOpts, amountA, amountB, minAmountA, minAmountB)
}

// AddKctLiquidityWithLimit is a paid mutator transaction binding the contract method 0xd9f25e88.
//
// Solidity: function addKctLiquidityWithLimit(uint256 amountA, uint256 amountB, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageTransactorSession) AddKctLiquidityWithLimit(amountA *big.Int, amountB *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKctLiquidityWithLimit(&_Storage.TransactOpts, amountA, amountB, minAmountA, minAmountB)
}

// AddKlayLiquidity is a paid mutator transaction binding the contract method 0x36e4245a.
//
// Solidity: function addKlayLiquidity(uint256 amount) payable returns()
func (_Storage *StorageTransactor) AddKlayLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addKlayLiquidity", amount)
}

// AddKlayLiquidity is a paid mutator transaction binding the contract method 0x36e4245a.
//
// Solidity: function addKlayLiquidity(uint256 amount) payable returns()
func (_Storage *StorageSession) AddKlayLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKlayLiquidity(&_Storage.TransactOpts, amount)
}

// AddKlayLiquidity is a paid mutator transaction binding the contract method 0x36e4245a.
//
// Solidity: function addKlayLiquidity(uint256 amount) payable returns()
func (_Storage *StorageTransactorSession) AddKlayLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKlayLiquidity(&_Storage.TransactOpts, amount)
}

// AddKlayLiquidityWithLimit is a paid mutator transaction binding the contract method 0x1b92cd07.
//
// Solidity: function addKlayLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) payable returns()
func (_Storage *StorageTransactor) AddKlayLiquidityWithLimit(opts *bind.TransactOpts, amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addKlayLiquidityWithLimit", amount, minAmountA, minAmountB)
}

// AddKlayLiquidityWithLimit is a paid mutator transaction binding the contract method 0x1b92cd07.
//
// Solidity: function addKlayLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) payable returns()
func (_Storage *StorageSession) AddKlayLiquidityWithLimit(amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKlayLiquidityWithLimit(&_Storage.TransactOpts, amount, minAmountA, minAmountB)
}

// AddKlayLiquidityWithLimit is a paid mutator transaction binding the contract method 0x1b92cd07.
//
// Solidity: function addKlayLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) payable returns()
func (_Storage *StorageTransactorSession) AddKlayLiquidityWithLimit(amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddKlayLiquidityWithLimit(&_Storage.TransactOpts, amount, minAmountA, minAmountB)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Storage *StorageTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Storage *StorageSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Storage *StorageTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Approve(&_Storage.TransactOpts, _spender, _value)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_Storage *StorageTransactor) ChangeFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeFee", _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_Storage *StorageSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeFee(&_Storage.TransactOpts, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_Storage *StorageTransactorSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeFee(&_Storage.TransactOpts, _fee)
}

// ChangeMiningRate is a paid mutator transaction binding the contract method 0x3d9222fe.
//
// Solidity: function changeMiningRate(uint256 _mining) returns()
func (_Storage *StorageTransactor) ChangeMiningRate(opts *bind.TransactOpts, _mining *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeMiningRate", _mining)
}

// ChangeMiningRate is a paid mutator transaction binding the contract method 0x3d9222fe.
//
// Solidity: function changeMiningRate(uint256 _mining) returns()
func (_Storage *StorageSession) ChangeMiningRate(_mining *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeMiningRate(&_Storage.TransactOpts, _mining)
}

// ChangeMiningRate is a paid mutator transaction binding the contract method 0x3d9222fe.
//
// Solidity: function changeMiningRate(uint256 _mining) returns()
func (_Storage *StorageTransactorSession) ChangeMiningRate(_mining *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeMiningRate(&_Storage.TransactOpts, _mining)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Storage *StorageTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Storage *StorageSession) ClaimReward() (*types.Transaction, error) {
	return _Storage.Contract.ClaimReward(&_Storage.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Storage *StorageTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Storage.Contract.ClaimReward(&_Storage.TransactOpts)
}

// ExchangeNeg is a paid mutator transaction binding the contract method 0x6029746c.
//
// Solidity: function exchangeNeg(address token, uint256 amount) returns(uint256)
func (_Storage *StorageTransactor) ExchangeNeg(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exchangeNeg", token, amount)
}

// ExchangeNeg is a paid mutator transaction binding the contract method 0x6029746c.
//
// Solidity: function exchangeNeg(address token, uint256 amount) returns(uint256)
func (_Storage *StorageSession) ExchangeNeg(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ExchangeNeg(&_Storage.TransactOpts, token, amount)
}

// ExchangeNeg is a paid mutator transaction binding the contract method 0x6029746c.
//
// Solidity: function exchangeNeg(address token, uint256 amount) returns(uint256)
func (_Storage *StorageTransactorSession) ExchangeNeg(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ExchangeNeg(&_Storage.TransactOpts, token, amount)
}

// ExchangePos is a paid mutator transaction binding the contract method 0xd013070c.
//
// Solidity: function exchangePos(address token, uint256 amount) returns(uint256)
func (_Storage *StorageTransactor) ExchangePos(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exchangePos", token, amount)
}

// ExchangePos is a paid mutator transaction binding the contract method 0xd013070c.
//
// Solidity: function exchangePos(address token, uint256 amount) returns(uint256)
func (_Storage *StorageSession) ExchangePos(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ExchangePos(&_Storage.TransactOpts, token, amount)
}

// ExchangePos is a paid mutator transaction binding the contract method 0xd013070c.
//
// Solidity: function exchangePos(address token, uint256 amount) returns(uint256)
func (_Storage *StorageTransactorSession) ExchangePos(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ExchangePos(&_Storage.TransactOpts, token, amount)
}

// GrabKlayFromFactory is a paid mutator transaction binding the contract method 0x1d68f84b.
//
// Solidity: function grabKlayFromFactory() payable returns()
func (_Storage *StorageTransactor) GrabKlayFromFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "grabKlayFromFactory")
}

// GrabKlayFromFactory is a paid mutator transaction binding the contract method 0x1d68f84b.
//
// Solidity: function grabKlayFromFactory() payable returns()
func (_Storage *StorageSession) GrabKlayFromFactory() (*types.Transaction, error) {
	return _Storage.Contract.GrabKlayFromFactory(&_Storage.TransactOpts)
}

// GrabKlayFromFactory is a paid mutator transaction binding the contract method 0x1d68f84b.
//
// Solidity: function grabKlayFromFactory() payable returns()
func (_Storage *StorageTransactorSession) GrabKlayFromFactory() (*types.Transaction, error) {
	return _Storage.Contract.GrabKlayFromFactory(&_Storage.TransactOpts)
}

// InitPool is a paid mutator transaction binding the contract method 0xf7d372f1.
//
// Solidity: function initPool(address to) returns()
func (_Storage *StorageTransactor) InitPool(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initPool", to)
}

// InitPool is a paid mutator transaction binding the contract method 0xf7d372f1.
//
// Solidity: function initPool(address to) returns()
func (_Storage *StorageSession) InitPool(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.InitPool(&_Storage.TransactOpts, to)
}

// InitPool is a paid mutator transaction binding the contract method 0xf7d372f1.
//
// Solidity: function initPool(address to) returns()
func (_Storage *StorageTransactorSession) InitPool(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.InitPool(&_Storage.TransactOpts, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Storage *StorageTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidity", amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Storage *StorageSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Storage *StorageTransactorSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, amount)
}

// RemoveLiquidityWithLimit is a paid mutator transaction binding the contract method 0xfde7d515.
//
// Solidity: function removeLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageTransactor) RemoveLiquidityWithLimit(opts *bind.TransactOpts, amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidityWithLimit", amount, minAmountA, minAmountB)
}

// RemoveLiquidityWithLimit is a paid mutator transaction binding the contract method 0xfde7d515.
//
// Solidity: function removeLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageSession) RemoveLiquidityWithLimit(amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityWithLimit(&_Storage.TransactOpts, amount, minAmountA, minAmountB)
}

// RemoveLiquidityWithLimit is a paid mutator transaction binding the contract method 0xfde7d515.
//
// Solidity: function removeLiquidityWithLimit(uint256 amount, uint256 minAmountA, uint256 minAmountB) returns()
func (_Storage *StorageTransactorSession) RemoveLiquidityWithLimit(amount *big.Int, minAmountA *big.Int, minAmountB *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityWithLimit(&_Storage.TransactOpts, amount, minAmountA, minAmountB)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Storage *StorageSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Storage *StorageTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Storage *StorageTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Storage *StorageSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Storage *StorageTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferFrom(&_Storage.TransactOpts, _from, _to, _value)
}

// UpdateMiningIndex is a paid mutator transaction binding the contract method 0xc1fac1d8.
//
// Solidity: function updateMiningIndex() returns()
func (_Storage *StorageTransactor) UpdateMiningIndex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "updateMiningIndex")
}

// UpdateMiningIndex is a paid mutator transaction binding the contract method 0xc1fac1d8.
//
// Solidity: function updateMiningIndex() returns()
func (_Storage *StorageSession) UpdateMiningIndex() (*types.Transaction, error) {
	return _Storage.Contract.UpdateMiningIndex(&_Storage.TransactOpts)
}

// UpdateMiningIndex is a paid mutator transaction binding the contract method 0xc1fac1d8.
//
// Solidity: function updateMiningIndex() returns()
func (_Storage *StorageTransactorSession) UpdateMiningIndex() (*types.Transaction, error) {
	return _Storage.Contract.UpdateMiningIndex(&_Storage.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}

// StorageAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Storage contract.
type StorageAddLiquidityIterator struct {
	Event *StorageAddLiquidity // Event containing the contract specifics and raw log

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
func (it *StorageAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddLiquidity)
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
		it.Event = new(StorageAddLiquidity)
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
func (it *StorageAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddLiquidity represents a AddLiquidity event raised by the Storage contract.
type StorageAddLiquidity struct {
	User      common.Address
	TokenA    common.Address
	AmountA   *big.Int
	TokenB    common.Address
	AmountB   *big.Int
	Liquidity *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xcfc390ea13b09661b205e48ee692d38f6798b0e0992a868008f841929b0e03d2.
//
// Solidity: event AddLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*StorageAddLiquidityIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &StorageAddLiquidityIterator{contract: _Storage.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xcfc390ea13b09661b205e48ee692d38f6798b0e0992a868008f841929b0e03d2.
//
// Solidity: event AddLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *StorageAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddLiquidity)
				if err := _Storage.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xcfc390ea13b09661b205e48ee692d38f6798b0e0992a868008f841929b0e03d2.
//
// Solidity: event AddLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) ParseAddLiquidity(log types.Log) (*StorageAddLiquidity, error) {
	event := new(StorageAddLiquidity)
	if err := _Storage.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// StorageChangeFeeIterator is returned from FilterChangeFee and is used to iterate over the raw logs and unpacked data for ChangeFee events raised by the Storage contract.
type StorageChangeFeeIterator struct {
	Event *StorageChangeFee // Event containing the contract specifics and raw log

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
func (it *StorageChangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageChangeFee)
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
		it.Event = new(StorageChangeFee)
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
func (it *StorageChangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageChangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageChangeFee represents a ChangeFee event raised by the Storage contract.
type StorageChangeFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeFee is a free log retrieval operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 _fee)
func (_Storage *StorageFilterer) FilterChangeFee(opts *bind.FilterOpts) (*StorageChangeFeeIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return &StorageChangeFeeIterator{contract: _Storage.contract, event: "ChangeFee", logs: logs, sub: sub}, nil
}

// WatchChangeFee is a free log subscription operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 _fee)
func (_Storage *StorageFilterer) WatchChangeFee(opts *bind.WatchOpts, sink chan<- *StorageChangeFee) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageChangeFee)
				if err := _Storage.contract.UnpackLog(event, "ChangeFee", log); err != nil {
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

// ParseChangeFee is a log parse operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 _fee)
func (_Storage *StorageFilterer) ParseChangeFee(log types.Log) (*StorageChangeFee, error) {
	event := new(StorageChangeFee)
	if err := _Storage.contract.UnpackLog(event, "ChangeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageChangeMiningRateIterator is returned from FilterChangeMiningRate and is used to iterate over the raw logs and unpacked data for ChangeMiningRate events raised by the Storage contract.
type StorageChangeMiningRateIterator struct {
	Event *StorageChangeMiningRate // Event containing the contract specifics and raw log

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
func (it *StorageChangeMiningRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageChangeMiningRate)
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
		it.Event = new(StorageChangeMiningRate)
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
func (it *StorageChangeMiningRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageChangeMiningRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageChangeMiningRate represents a ChangeMiningRate event raised by the Storage contract.
type StorageChangeMiningRate struct {
	Mining *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeMiningRate is a free log retrieval operation binding the contract event 0xbbe8b2527b3f8a4889631c97b9a60cdf1affe7ae3f6216a8a4547e7d99021446.
//
// Solidity: event ChangeMiningRate(uint256 _mining)
func (_Storage *StorageFilterer) FilterChangeMiningRate(opts *bind.FilterOpts) (*StorageChangeMiningRateIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ChangeMiningRate")
	if err != nil {
		return nil, err
	}
	return &StorageChangeMiningRateIterator{contract: _Storage.contract, event: "ChangeMiningRate", logs: logs, sub: sub}, nil
}

// WatchChangeMiningRate is a free log subscription operation binding the contract event 0xbbe8b2527b3f8a4889631c97b9a60cdf1affe7ae3f6216a8a4547e7d99021446.
//
// Solidity: event ChangeMiningRate(uint256 _mining)
func (_Storage *StorageFilterer) WatchChangeMiningRate(opts *bind.WatchOpts, sink chan<- *StorageChangeMiningRate) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ChangeMiningRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageChangeMiningRate)
				if err := _Storage.contract.UnpackLog(event, "ChangeMiningRate", log); err != nil {
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

// ParseChangeMiningRate is a log parse operation binding the contract event 0xbbe8b2527b3f8a4889631c97b9a60cdf1affe7ae3f6216a8a4547e7d99021446.
//
// Solidity: event ChangeMiningRate(uint256 _mining)
func (_Storage *StorageFilterer) ParseChangeMiningRate(log types.Log) (*StorageChangeMiningRate, error) {
	event := new(StorageChangeMiningRate)
	if err := _Storage.contract.UnpackLog(event, "ChangeMiningRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageExchangeNegIterator is returned from FilterExchangeNeg and is used to iterate over the raw logs and unpacked data for ExchangeNeg events raised by the Storage contract.
type StorageExchangeNegIterator struct {
	Event *StorageExchangeNeg // Event containing the contract specifics and raw log

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
func (it *StorageExchangeNegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageExchangeNeg)
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
		it.Event = new(StorageExchangeNeg)
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
func (it *StorageExchangeNegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageExchangeNegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageExchangeNeg represents a ExchangeNeg event raised by the Storage contract.
type StorageExchangeNeg struct {
	TokenA  common.Address
	AmountA *big.Int
	TokenB  common.Address
	AmountB *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExchangeNeg is a free log retrieval operation binding the contract event 0x941b04a0d2854767c6f795fa84b0a74f4f95c1fffac9cbdf66d7b3fddba60dfd.
//
// Solidity: event ExchangeNeg(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) FilterExchangeNeg(opts *bind.FilterOpts) (*StorageExchangeNegIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ExchangeNeg")
	if err != nil {
		return nil, err
	}
	return &StorageExchangeNegIterator{contract: _Storage.contract, event: "ExchangeNeg", logs: logs, sub: sub}, nil
}

// WatchExchangeNeg is a free log subscription operation binding the contract event 0x941b04a0d2854767c6f795fa84b0a74f4f95c1fffac9cbdf66d7b3fddba60dfd.
//
// Solidity: event ExchangeNeg(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) WatchExchangeNeg(opts *bind.WatchOpts, sink chan<- *StorageExchangeNeg) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ExchangeNeg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageExchangeNeg)
				if err := _Storage.contract.UnpackLog(event, "ExchangeNeg", log); err != nil {
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

// ParseExchangeNeg is a log parse operation binding the contract event 0x941b04a0d2854767c6f795fa84b0a74f4f95c1fffac9cbdf66d7b3fddba60dfd.
//
// Solidity: event ExchangeNeg(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) ParseExchangeNeg(log types.Log) (*StorageExchangeNeg, error) {
	event := new(StorageExchangeNeg)
	if err := _Storage.contract.UnpackLog(event, "ExchangeNeg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageExchangePosIterator is returned from FilterExchangePos and is used to iterate over the raw logs and unpacked data for ExchangePos events raised by the Storage contract.
type StorageExchangePosIterator struct {
	Event *StorageExchangePos // Event containing the contract specifics and raw log

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
func (it *StorageExchangePosIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageExchangePos)
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
		it.Event = new(StorageExchangePos)
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
func (it *StorageExchangePosIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageExchangePosIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageExchangePos represents a ExchangePos event raised by the Storage contract.
type StorageExchangePos struct {
	TokenA  common.Address
	AmountA *big.Int
	TokenB  common.Address
	AmountB *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExchangePos is a free log retrieval operation binding the contract event 0x022d176d604c15661a2acf52f28fd69bdd2c755884c08a67132ffeb8098330e0.
//
// Solidity: event ExchangePos(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) FilterExchangePos(opts *bind.FilterOpts) (*StorageExchangePosIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ExchangePos")
	if err != nil {
		return nil, err
	}
	return &StorageExchangePosIterator{contract: _Storage.contract, event: "ExchangePos", logs: logs, sub: sub}, nil
}

// WatchExchangePos is a free log subscription operation binding the contract event 0x022d176d604c15661a2acf52f28fd69bdd2c755884c08a67132ffeb8098330e0.
//
// Solidity: event ExchangePos(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) WatchExchangePos(opts *bind.WatchOpts, sink chan<- *StorageExchangePos) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ExchangePos")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageExchangePos)
				if err := _Storage.contract.UnpackLog(event, "ExchangePos", log); err != nil {
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

// ParseExchangePos is a log parse operation binding the contract event 0x022d176d604c15661a2acf52f28fd69bdd2c755884c08a67132ffeb8098330e0.
//
// Solidity: event ExchangePos(address tokenA, uint256 amountA, address tokenB, uint256 amountB)
func (_Storage *StorageFilterer) ParseExchangePos(log types.Log) (*StorageExchangePos, error) {
	event := new(StorageExchangePos)
	if err := _Storage.contract.UnpackLog(event, "ExchangePos", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageGiveRewardIterator is returned from FilterGiveReward and is used to iterate over the raw logs and unpacked data for GiveReward events raised by the Storage contract.
type StorageGiveRewardIterator struct {
	Event *StorageGiveReward // Event containing the contract specifics and raw log

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
func (it *StorageGiveRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageGiveReward)
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
		it.Event = new(StorageGiveReward)
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
func (it *StorageGiveRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageGiveRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageGiveReward represents a GiveReward event raised by the Storage contract.
type StorageGiveReward struct {
	User      common.Address
	Amount    *big.Int
	LastIndex *big.Int
	RewardSum *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGiveReward is a free log retrieval operation binding the contract event 0x5665ebb50e1da7888dabf11997217c5f8d9019c1d10dc7fed8a300556e7c5de4.
//
// Solidity: event GiveReward(address user, uint256 amount, uint256 lastIndex, uint256 rewardSum)
func (_Storage *StorageFilterer) FilterGiveReward(opts *bind.FilterOpts) (*StorageGiveRewardIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "GiveReward")
	if err != nil {
		return nil, err
	}
	return &StorageGiveRewardIterator{contract: _Storage.contract, event: "GiveReward", logs: logs, sub: sub}, nil
}

// WatchGiveReward is a free log subscription operation binding the contract event 0x5665ebb50e1da7888dabf11997217c5f8d9019c1d10dc7fed8a300556e7c5de4.
//
// Solidity: event GiveReward(address user, uint256 amount, uint256 lastIndex, uint256 rewardSum)
func (_Storage *StorageFilterer) WatchGiveReward(opts *bind.WatchOpts, sink chan<- *StorageGiveReward) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "GiveReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageGiveReward)
				if err := _Storage.contract.UnpackLog(event, "GiveReward", log); err != nil {
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

// ParseGiveReward is a log parse operation binding the contract event 0x5665ebb50e1da7888dabf11997217c5f8d9019c1d10dc7fed8a300556e7c5de4.
//
// Solidity: event GiveReward(address user, uint256 amount, uint256 lastIndex, uint256 rewardSum)
func (_Storage *StorageFilterer) ParseGiveReward(log types.Log) (*StorageGiveReward, error) {
	event := new(StorageGiveReward)
	if err := _Storage.contract.UnpackLog(event, "GiveReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Storage contract.
type StorageRemoveLiquidityIterator struct {
	Event *StorageRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *StorageRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRemoveLiquidity)
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
		it.Event = new(StorageRemoveLiquidity)
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
func (it *StorageRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRemoveLiquidity represents a RemoveLiquidity event raised by the Storage contract.
type StorageRemoveLiquidity struct {
	User      common.Address
	TokenA    common.Address
	AmountA   *big.Int
	TokenB    common.Address
	AmountB   *big.Int
	Liquidity *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x31a83f5720229e02936f2f1b2260a050bbd8377dce85422763735b783106deb1.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*StorageRemoveLiquidityIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &StorageRemoveLiquidityIterator{contract: _Storage.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x31a83f5720229e02936f2f1b2260a050bbd8377dce85422763735b783106deb1.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *StorageRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRemoveLiquidity)
				if err := _Storage.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x31a83f5720229e02936f2f1b2260a050bbd8377dce85422763735b783106deb1.
//
// Solidity: event RemoveLiquidity(address user, address tokenA, uint256 amountA, address tokenB, uint256 amountB, uint256 liquidity)
func (_Storage *StorageFilterer) ParseRemoveLiquidity(log types.Log) (*StorageRemoveLiquidity, error) {
	event := new(StorageRemoveLiquidity)
	if err := _Storage.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// StorageUpdateMiningIndexIterator is returned from FilterUpdateMiningIndex and is used to iterate over the raw logs and unpacked data for UpdateMiningIndex events raised by the Storage contract.
type StorageUpdateMiningIndexIterator struct {
	Event *StorageUpdateMiningIndex // Event containing the contract specifics and raw log

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
func (it *StorageUpdateMiningIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpdateMiningIndex)
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
		it.Event = new(StorageUpdateMiningIndex)
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
func (it *StorageUpdateMiningIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpdateMiningIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpdateMiningIndex represents a UpdateMiningIndex event raised by the Storage contract.
type StorageUpdateMiningIndex struct {
	LastMined   *big.Int
	MiningIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateMiningIndex is a free log retrieval operation binding the contract event 0x08958c674d1654780d97078fb3a6d043c62f4bb7d8c15f9dc779f8409a856f04.
//
// Solidity: event UpdateMiningIndex(uint256 lastMined, uint256 miningIndex)
func (_Storage *StorageFilterer) FilterUpdateMiningIndex(opts *bind.FilterOpts) (*StorageUpdateMiningIndexIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UpdateMiningIndex")
	if err != nil {
		return nil, err
	}
	return &StorageUpdateMiningIndexIterator{contract: _Storage.contract, event: "UpdateMiningIndex", logs: logs, sub: sub}, nil
}

// WatchUpdateMiningIndex is a free log subscription operation binding the contract event 0x08958c674d1654780d97078fb3a6d043c62f4bb7d8c15f9dc779f8409a856f04.
//
// Solidity: event UpdateMiningIndex(uint256 lastMined, uint256 miningIndex)
func (_Storage *StorageFilterer) WatchUpdateMiningIndex(opts *bind.WatchOpts, sink chan<- *StorageUpdateMiningIndex) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UpdateMiningIndex")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpdateMiningIndex)
				if err := _Storage.contract.UnpackLog(event, "UpdateMiningIndex", log); err != nil {
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

// ParseUpdateMiningIndex is a log parse operation binding the contract event 0x08958c674d1654780d97078fb3a6d043c62f4bb7d8c15f9dc779f8409a856f04.
//
// Solidity: event UpdateMiningIndex(uint256 lastMined, uint256 miningIndex)
func (_Storage *StorageFilterer) ParseUpdateMiningIndex(log types.Log) (*StorageUpdateMiningIndex, error) {
	event := new(StorageUpdateMiningIndex)
	if err := _Storage.contract.UnpackLog(event, "UpdateMiningIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
