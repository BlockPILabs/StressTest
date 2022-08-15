package main

import (
	"flag"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var url string
var loglevel string
var vuser int
var vtime time.Duration
var factoryAbiData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"teamWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExchangeImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newExImp\",\"type\":\"address\"}],\"name\":\"_setExchangeImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"halfLife\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"teamAward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"teamRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newImp\",\"type\":\"address\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unfreezeBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minableBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"},{\"name\":\"_exchangeImplementation\",\"type\":\"address\"},{\"name\":\"_miningAmount\",\"type\":\"uint256\"},{\"name\":\"_halfLife\",\"type\":\"uint256\"},{\"name\":\"_minableBlock\",\"type\":\"uint256\"},{\"name\":\"_unfreezeBlock\",\"type\":\"uint256\"},{\"name\":\"_teamRatio\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}
var ercAbiData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_version\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetTokenInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}
var pairAbiData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"addKlayLiquidityWithLimit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"grabKlayFromFactory\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addKlayLiquidity\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mining\",\"type\":\"uint256\"}],\"name\":\"changeMiningRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeNeg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mining\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amountA\",\"type\":\"uint256\"},{\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"addKctLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastMined\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"estimateNeg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateMiningIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangePos\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amountA\",\"type\":\"uint256\"},{\"name\":\"amountB\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"addKctLiquidityWithLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"estimatePos\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLastIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"initPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"minAmountA\",\"type\":\"uint256\"},{\"name\":\"minAmountB\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityWithLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_mining\",\"type\":\"uint256\"}],\"name\":\"ChangeMiningRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lastMined\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"miningIndex\",\"type\":\"uint256\"}],\"name\":\"UpdateMiningIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardSum\",\"type\":\"uint256\"}],\"name\":\"GiveReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"ExchangePos\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"ExchangeNeg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

var factoryAbi = factoryAbiData.ABI
var ercAbi = ercAbiData.ABI
var pairAbi = pairAbiData.ABI

type swapToken struct {
	coinA common.Address
	coinB common.Address
}

var factoryAddr = common.HexToAddress("0xc6a2ad8cc6e4a7e08fc37cc5954be07d499e7654")
var swapTokenLists = []swapToken{}

// initFlag 设置程序命令行参数
func initFlag() {
	flag.IntVar(&vuser, "u", 10, "Number of virtual users running")
	flag.DurationVar(&vtime, "d", 5*time.Second, "running time")
	flag.StringVar(&loglevel, "l", "error", "set logLevel debug info warn error")
	flag.StringVar(&url, "url", "http://156.146.44.2:31271", "set request interface")
	flag.Parse()
}

// 初始化全局变量
func initVer() {
	// 设置时区
	time.Local = time.FixedZone("CST", 8*3600)

	swapTokenLists = append(swapTokenLists,
		swapToken{oUSDT, oETH},
		swapToken{oUSDT, SIG},
		swapToken{oUSDT, KONE},
		swapToken{oUSDT, KPG},
		swapToken{oUSDT, WEMIX},
		swapToken{oUSDT, KP},
		swapToken{oUSDT, KAI},
		swapToken{oUSDT, AOA},
		swapToken{oUSDT, CLAM},
		swapToken{oUSDT, TOX},
		swapToken{oUSDT, ZEMIT},
		swapToken{oUSDT, Q2},
		swapToken{oUSDT, oXRP},
		swapToken{oUSDT, CEIK},
		swapToken{oUSDT, WALK},
		swapToken{oUSDT, RSTAR},
		swapToken{oUSDT, bus},
		swapToken{oUSDT, PER},
		swapToken{oUSDT, MOOI},
		swapToken{oUSDT, EKL},
		swapToken{oUSDT, MON},
		swapToken{oUSDT, NGIT},
		swapToken{oUSDT, KSTAR},
		swapToken{oUSDT, test_h},
		swapToken{oUSDT, ORB},
		swapToken{oUSDT, JUN},
		swapToken{oUSDT, IJM},
		swapToken{oUSDT, oORC},
		swapToken{oUSDT, nice},
		swapToken{oUSDT, oUSDC},
		swapToken{oUSDT, USDF},
		swapToken{oUSDT, NoA},
		swapToken{oUSDT, sBWPM},
		swapToken{oUSDT, CLA},
		swapToken{oUSDT, KSP},
		swapToken{KSP, oUSDT},
		swapToken{KSP, oETH},
		swapToken{KSP, FINIX},
		swapToken{KSP, NolzeBoost},
		swapToken{KSP, sKAI},
		swapToken{KSP, NGIT},
		swapToken{KSP, oPACE},
		swapToken{KSP, NIK},
		swapToken{KSP, CX2},
		swapToken{KSP, STAR},
		swapToken{KSP, HALF},
		swapToken{KSP, oMESH},
		swapToken{KSP, AOA},
		swapToken{KSP, oWBTC},
		swapToken{KSP, CLBK},
		swapToken{KSP, PUNK},
		swapToken{KSP, oORC},
		swapToken{KSP, Nolze},
		swapToken{KSP, oXRP},
		swapToken{KSP, bus},
		swapToken{KSP, KDAI},
		swapToken{KSP, PER},
		swapToken{KSP, oUSDC},
		swapToken{KSP, AZIT},
		swapToken{KSP, KFI},
		swapToken{KSP, oXDT},
		swapToken{KSP, MIX},
		swapToken{KSP, KLEVA},
		swapToken{KSP, MOOI},
		swapToken{KSP, busB},
		swapToken{KSP, CEIK},
		swapToken{KSP, MBX},
		swapToken{KSP, MKC},
		swapToken{KSP, SPRX},
		swapToken{KSP, ABC},
		swapToken{KSP, IBP},
		swapToken{KSP, COM},
		swapToken{KSP, GHUB},
		swapToken{KSP, INS},
		swapToken{KSP, ZEMIT},
		swapToken{KSP, oBNB},
	)
}

var KDAI = common.HexToAddress("0x5c74070fdea071359b86082bd9f9b3deaafbe32b")
var SKLAY = common.HexToAddress("0xa323d7386b671e8799dca3582d6658fdcdcd940a")
var oETH = common.HexToAddress("0x34d21b1e550d73cee41151c77f3c73359527a396")
var oWBTC = common.HexToAddress("0x16d0e1fbd024c600ca0380a4c5d57ee7a2ecbf9c")
var oUSDT = common.HexToAddress("0xcee8faf64bb97a73bb51e115aa89c17ffa8dd167")
var oORC = common.HexToAddress("0xfe41102f325deaa9f303fdd9484eb5911a7ba557")
var ISR = common.HexToAddress("0x9657fb399847d85a9c1a234ece9ca09d5c00f466")
var WIKEN = common.HexToAddress("0x275f942985503d8ce9558f8377cc526a3aba3566")
var SSX = common.HexToAddress("0xdcd62c57182e780e23d2313c4782709da85b9d6c")
var ABL = common.HexToAddress("0x46f307b58bf05ff089ba23799fae0e518557f87c")
var REDi = common.HexToAddress("0x1cd3828a2b62648dbe98d6f5748a6b1df08ac7bb")
var KSP = common.HexToAddress("0xc6a2ad8cc6e4a7e08fc37cc5954be07d499e7654")
var AGOV = common.HexToAddress("0x588c62ed9aa7367d7cd9c2a9aaac77e44fe8221b")
var HINT = common.HexToAddress("0x4dd402a7d54eaa8147cb6ff252afe5be742bdf40")
var BEE = common.HexToAddress("0xa128e62cfb454ab5b580a7385de2f228ad7b69d1")
var WEMIX = common.HexToAddress("0x5096db80b21ef45230c9e423c373f1fc9c0198dd")
var BBC = common.HexToAddress("0x321bc0b63efb1e4af08ec6d20c85d5e94ddaaa18")
var oTRIX = common.HexToAddress("0x0c1d7ce4982fd63b1bc77044be1da05c995e4463")
var PIB = common.HexToAddress("0xafde910130c335fa5bd5fe991053e3e0a49dce7b")
var CLBK = common.HexToAddress("0xc4407f7dc4b37275c9ce0f839652b393e13ff3d1")
var KSLO = common.HexToAddress("0xf7689072554d1e85fa9d96151f528764277d7db2")
var KSVE = common.HexToAddress("0x065a9ddbbdd48c4189984e2f7aeda3834bd1eac7")
var oHANDY = common.HexToAddress("0x3f34671fba493ab39fbf4ecac2943ee62b654a88")
var MNR = common.HexToAddress("0x27dcd181459bcddc63c37bab1e404a313c0dfd79")
var TRCL = common.HexToAddress("0x4b91c67a89d4c4b2a4ed9fcde6130d7495330972")
var oXRP = common.HexToAddress("0x9eaefb09fe4aabfbe6b1ca316a3c36afc83a393f")
var KDUCATO = common.HexToAddress("0x91e0d7b228a33072d9b3209cf507f78a4bd835f2")
var TEMCO = common.HexToAddress("0x3b3b30a76d169f72a0a38ae01b0d6e0fbee3cc2e")
var HIBS = common.HexToAddress("0xe06b40df899b9717b4e6b50711e1dc72d08184cf")
var oBELT = common.HexToAddress("0xdfe180e288158231ffa5faf183eca3301344a51f")
var oBNB = common.HexToAddress("0x574e9c26bda8b95d7329505b4657103710eb32ea")
var oQBZ = common.HexToAddress("0x507efa4e365fd5def42cb05ae3ecb51a30321588")
var oSTPL = common.HexToAddress("0x49a767b188b9d782d7b0efcd485ca3796390198e")
var oAUTO = common.HexToAddress("0x8583063110b5d29036eced4db1cc147e78a86a77")
var oTON = common.HexToAddress("0x100bc15ae8b489c771d9740ea0bb1aea945a1f67")
var BUZ = common.HexToAddress("0x75ad14d0360408dc6f8163e5dfb51aad516f4afd")
var SIX = common.HexToAddress("0xef82b1c6a550e730d8283e1edd4977cd01faf435")
var KAI = common.HexToAddress("0xe950bdcfa4d1e45472e76cf967db93dbfc51ba3e")
var sKAI = common.HexToAddress("0x37d46c6813b121d6a27ed263aef782081ae95434")
var oRAI = common.HexToAddress("0xb40178be0fcf89d0051682e5512a8bab56b9ec3e")
var TURK = common.HexToAddress("0x8c783809332be7734fa782eb5139861721f77b33")
var oRUSH = common.HexToAddress("0x2fade69ba4dcb112c530c48fdf41fc071685cede")
var INFR = common.HexToAddress("0x12d8a34e68f8c0333f10265a6930381657a9cc18")
var oISDT = common.HexToAddress("0x1c1187ff22bb50a2cdcb1e1d683fcb5e8a42915f")
var sBWPM = common.HexToAddress("0xf4546e1d3ad590a3c6d178d671b3bc0e8a81e27d")
var PER = common.HexToAddress("0x7eee60a000986e9efe7f5c90340738558c24317b")
var KFI = common.HexToAddress("0xdb116e2dc96b4e69e3544f41b50550436579979a")
var KICX = common.HexToAddress("0x8ef60f0a5a2db984431934f8659058e87cd5c70a")
var bKAI = common.HexToAddress("0x968d93a44b3ef61168ca621a59ddc0e56583e592")
var TESTTTT = common.HexToAddress("0xcab47614b96b85538f2423bda52342bdca7742b6")
var TESTAAA = common.HexToAddress("0x6dd78dc9fcb235284e7a02cab57ea4dd06f568e4")
var HOUSE = common.HexToAddress("0x158beff8c8cdebd64654add5f6a1d9937e73536c")
var AKLAY = common.HexToAddress("0x74ba03198fed2b15a51af242b9c63faf3c8f4d34")
var ORCA = common.HexToAddress("0xa4547080f3310b9ec4ed4b08fbc3762c6d294cc9")
var vKAI = common.HexToAddress("0x44efe1ec288470276e29ac3adb632bff990e2e1f")
var oTALK = common.HexToAddress("0xa6e559f2954c7e2c017bbbea85d8513f34d03d75")
var FINIX = common.HexToAddress("0xd51c337147c8033a43f3b5ce0023382320c113aa")
var DONT1BUY = common.HexToAddress("0x6499b23434d61a31c9da5af808fc0b9ee0c76b55")
var JUN = common.HexToAddress("0x5a55a1cd5cc5e89019300213f9faf20f57361d43")
var JUNS = common.HexToAddress("0x69df45d36341f6bad3c4beffb9e77f2b74709c40")
var IMI = common.HexToAddress("0xdde2154f47e80c8721c2efbe02834ae056284368")
var WOOD = common.HexToAddress("0x7b7363cf78662b638a87f63871c302be363ddc7a")
var oAERGO = common.HexToAddress("0x292725810ab3dde0a01e19acd4e8e9d6c073773b")
var MM = common.HexToAddress("0x96035fbdd4cb888862ee28c9d8fdadef78311cc9")
var ksDUNAMU = common.HexToAddress("0x1add8ba5a695063962abe9b7da70ed2440006def")
var ksYANOLJA = common.HexToAddress("0xa5c1cd7b9a015243a0ff8f99c5dbe17647175af0")
var ksCOINBASE = common.HexToAddress("0x88bfd174f9076519a45979ce3122bc15883c0691")
var TESTTOKEN = common.HexToAddress("0x63d8693539533b86d92c9099997b3bd8bf451b14")
var TESTSTABLE = common.HexToAddress("0xcdc259f7ce752b3cadce3b4fea3016e7cdc23ebf")
var ksETH = common.HexToAddress("0xdfc05e7a28ed3a1c22bc7c22383764a4732ead23")
var ksSOL = common.HexToAddress("0x7a85836f66dbbd53f457855de243f5aa28051e33")
var ksLUNA = common.HexToAddress("0x4afe41cae3bd5133991384d642b82f2d0539e3aa")
var ICEAMERICANO = common.HexToAddress("0xfed6892d9bcbfe37fa021f6a019a224598c7f2ab")
var CLAM = common.HexToAddress("0xba9725eaccf07044625f1d232ef682216f5371c2")
var NWC = common.HexToAddress("0x36e936d5f4b6ab59f232da22ce53650dd80a4386")
var JSD = common.HexToAddress("0x76264ad1b50852c4d8efb55bfaf154dd5a80d0c2")
var XE = common.HexToAddress("0x3513b2bc58f1f260107fd1ee0dabb5b0637b9ed5")
var oUSDC = common.HexToAddress("0x754288077d0ff82af7a5317c7cb8c444d421d103")
var BYPE = common.HexToAddress("0x2ef5f2642674f768b4efe9a7de470a6a68bcb8f3")
var TOR = common.HexToAddress("0x1d8246a6e73473ce4e21bc7e40bd5c3cef7930d5")
var KOKOA = common.HexToAddress("0xb15183d0d4d5e86ba702ce9bb7b633376e7db29f")
var KSD = common.HexToAddress("0x4fa62f1f404188ce860c8f0041d6ac3765a72e67")
var MIX = common.HexToAddress("0xdd483a970a7a7fef2b223c3510fac852799a88bf")
var CT = common.HexToAddress("0x7f223b1607171b81ebd68d22f1ca79157fd4a44b")
var TURNUP = common.HexToAddress("0xd742b1a5af898bcb4b6aff5027e6ab9adee97412")
var oQBT = common.HexToAddress("0xfe4cd053f1e9200e784b7d60b54e6aa16e09406a")
var oDON = common.HexToAddress("0x5388ce775de8f7a69d17fd5caa9f7dbfee65dfce")
var ksAPPLE = common.HexToAddress("0x5844b02cc0ab5d5a18be7dde4e245f5edec449ce")
var ksTESLA = common.HexToAddress("0x1e3a300601aa95ab7ea39bb72c3272716ef1426b")
var ksFarm = common.HexToAddress("0xef8245fdc4087adef6e861140138c4443d8d6b5d")
var CEIK = common.HexToAddress("0x18814b01b5cc76f7043e10fd268cc4364df47da0")
var KSTA = common.HexToAddress("0xe7d3b78f032e70fabfdb8c0741ea74f775deb32d")
var Q2 = common.HexToAddress("0x06d566491c858f50ee412bc475ac29260650f43c")
var bus = common.HexToAddress("0xf445e3d0f88c4c2c8a2751180ae4a525789cfe32")
var oDFA = common.HexToAddress("0x11dc950ef29594cc19eb573811339df20c86c800")
var PICS = common.HexToAddress("0xf87a3cf2f1dc059019455323edafb2667ea5cbe9")
var oCBANK = common.HexToAddress("0x52f4c436c9aab5b5d0dd31fb3fb8f253fd6cb285")
var BORA = common.HexToAddress("0x02cbe46fb8a1f579254a9b485788f2d86cad51aa")
var KAPT = common.HexToAddress("0xddc42416f16176d835311f710ee78aff63705b37")
var STC = common.HexToAddress("0x113e52528e5848e6ceceb3d8a8c4bd689f793469")
var JUNB = common.HexToAddress("0x01cb8563e9c4703f4e6b9fa09edeaed0e5948f28")
var KDG = common.HexToAddress("0x02e7d9ad54a19a9a0721d9515cf9f80f9547d771")
var oBIOT = common.HexToAddress("0x8ff0586b6eea63a35e73d09237b4a58b3056f274")
var KRNO = common.HexToAddress("0xd676e57ca65b827feb112ad81ff738e7b6c1048d")
var META = common.HexToAddress("0xe815a060b9279eba642f8c889fab7afc0d0aca63")
var tAVELK = common.HexToAddress("0x0b430ad7bf84eb307e221f0e66216205502f835d")
var oTOM = common.HexToAddress("0x3886df0cb94f4bf7e0ac3dfe2d1db40d2ee2293b")
var aKAI = common.HexToAddress("0xb57e0038e8027c3de8126a07cac371f31c9c229e")
var OYAT = common.HexToAddress("0x143d71be70dc518a9e9c25b6008d9353b3698d26")
var MUDOL = common.HexToAddress("0x45dbbbcdff605af5fe27fd5e93b9f3f1bb25d429")
var oGALA = common.HexToAddress("0x2842a6d0c182e3f1cf4556311c48a7706d7ba6ad")
var oDOTR = common.HexToAddress("0x8bc28c926a0fe54b5c56a329cd3b129cc52ae099")
var MPWR = common.HexToAddress("0xad27ace6f0f6cef2c192a3c8f0f3fa2611154eb3")
var oXVS = common.HexToAddress("0x735106530578fb0227422de25bb32c9adfb5ea2e")
var oBUSD = common.HexToAddress("0x210bc03f49052169d5588a52c317f71cf2078b85")
var oCAKE = common.HexToAddress("0x36e5ea82a099e8188bd5af5709b23628076de822")
var DANTA = common.HexToAddress("0x30763e9a3898b9a76d0a541380d927a50b9bbe81")
var oCYCLUB = common.HexToAddress("0xf1ec6fc5b9f280ed43b45d2ba60874a77f343c60")
var PUNK = common.HexToAddress("0x37c38b19a6ba325486da87f946e72dc93e0ab39a")
var oTRV = common.HexToAddress("0xa23e07f7a61ce663b9ca6d6683acaaa28ec3070f")
var oPACE = common.HexToAddress("0xbc5d3fb02514f975060d35000e99c54253002bd4")
var CLA = common.HexToAddress("0xcf87f94fd8f6b6f0b479771f10df672f99eada63")
var STAR = common.HexToAddress("0xbd8df801b7516a088736342c82cae56687e8282b")
var oMTS = common.HexToAddress("0x4e8975380f04ec53adfdef8f5f8246c4e04291f7")
var KPAX = common.HexToAddress("0x95f04d09a8dc87edcf1ba6fed443993fa2466465")
var GCT = common.HexToAddress("0x0a7356f5df9179c977d4ae5d285809a60f4797e4")
var NOAH = common.HexToAddress("0x8aa6b6b3d6cf0b20c922b626d55e60c7ea706648")
var KLEVA = common.HexToAddress("0x5fff3a6c16c2208103f318f4713d4d90601a7313")
var TnT = common.HexToAddress("0x17d30e878ba5a546c76fada32d7a30c876fadb49")
var JUNC = common.HexToAddress("0xd675dae87d8740b2163b4e232ee51a880495e6c7")
var Ksh = common.HexToAddress("0xbfd4542524a426b26caa9ad11c1755b1087a5737")
var RaGa = common.HexToAddress("0xca03afda20444e4131cf3de5fb9cd7991b9874bd")
var GOLF = common.HexToAddress("0x83c3b5a9a9d1f1438f2505ba972366eecfc4488e")
var LAMP = common.HexToAddress("0x325aa5489d446d2c8cc3832a57d95d207107a4d1")
var MEPS = common.HexToAddress("0x1e7af33c36c5f942687b8126af2012ea90eea041")
var oSOBA = common.HexToAddress("0xc5523901990c8977e34802347e8e2d7fac0437eb")
var KONGZ = common.HexToAddress("0x2cacf8eb48f609d7dce06a4e4f9ff4c22563eb62")
var BANANA = common.HexToAddress("0xf6a6fc906b38eab8c9eb3a0ba9606262704af091")
var TBK = common.HexToAddress("0x995f6cde1753799bdb37c838ccfbdca2bc160df5")
var INS = common.HexToAddress("0xdf9e1b5a30d6175cabaaf39964dd979e84753eb1")
var CATC = common.HexToAddress("0xeddbbc44fe7cedaf0a2a7b40971a23dae82c1c8c")
var IBP = common.HexToAddress("0x43de991a0d9b666a215f3eb5801accd260092c2c")
var IJM = common.HexToAddress("0x0268dbed3832b87582b1fa508acf5958cbb1cd74")
var oDRC = common.HexToAddress("0x923ff99c3ca96c77e4c111d80b71d2504c364f7a")
var KREDIT = common.HexToAddress("0x24703f8497412912106761210bdc30c44ee61b2f")
var KP = common.HexToAddress("0x2b5d75db09af26e53d051155f5eae811db7aef67")
var KPs = common.HexToAddress("0xe1376ab327b6deb7bebaee1329eb94574d51a8d9")
var test_h = common.HexToAddress("0x03abbd3d662ff73e7a53d6d71b16436756acad9d")
var RaPP = common.HexToAddress("0xcb1b1309bb70e21a3838987d1cfdb3db885efaaa")
var UKIKI = common.HexToAddress("0x6903c62861c687e2e41c06a68193adfba5b97254")
var MKC = common.HexToAddress("0x119883ee408aa5b9625c5d09a79fa8be9f9f6017")
var TOP = common.HexToAddress("0x27cf1333e3e6ec8b518fb84a5ff7aea8cf790f3d")
var TOP2 = common.HexToAddress("0x6d2b550f3d86c378627478a8882ae9842471c266")
var AZIT = common.HexToAddress("0x6cef6dd9a3c4ad226b8b66effeea2c125df194f1")
var HOOK = common.HexToAddress("0x8ef50fa375fc64b9815e51f28f4b83c05d57ac43")
var BT = common.HexToAddress("0x406ec348311fc463e677ddd77c42ce28877dd303")
var BS = common.HexToAddress("0x35f9e17f1a1d0ab6c3e43be8680952f7bda5305f")
var COM = common.HexToAddress("0x29435457053d167a2b1f6f2d54d4176866ffb5f9")
var CRE = common.HexToAddress("0xde3b13c45d4c2c3f4965932314e23279f426239f")
var EKL = common.HexToAddress("0x807c4e063eb0ac21e8eef7623a6ed50a8ede58ca")
var KIDS = common.HexToAddress("0xe0f2a679390efb0507ae8f99db4b7832202ac808")
var NGIT = common.HexToAddress("0x46db825593ca7c3fdfc9ccb5850ea96c39b79330")
var ducato = common.HexToAddress("0xb7ab8c205dc282f83d0511b1dcd22a3ed4739597")
var SALT = common.HexToAddress("0x3247abb921c83f81b406e1a87fb7bfa6f79262d0")
var BTIS = common.HexToAddress("0xa31b5c5027e4b3f4e7b63e5bec2e598d8bf870c2")
var HC = common.HexToAddress("0xe43686e3a798741ea761cd8da6785e27b92ca623")
var NIK = common.HexToAddress("0xdabee145a1395e09280c23ea9aa71caca35a1ec0")
var RaB = common.HexToAddress("0xeec98f041a7325e65b9140107dfa68db263a2ac4")
var FDM = common.HexToAddress("0xb1834e4e773a180168f2292e036ca8e17f86196f")
var BOMUL = common.HexToAddress("0x5655ee0628ad3348cb7b60e8102680bb0d7f0de1")
var busB = common.HexToAddress("0x730c261202adc5f30cf9b892f665f35c095b81a6")
var RaM = common.HexToAddress("0x138630e76c0184674536d3503482ea4a785baf7b")
var Nolze = common.HexToAddress("0x3788871fd2c57a34335e5d7d1273e85f66e2fd5e")
var DLP = common.HexToAddress("0x4d87baa66061a3bb391456576cc49b77540b9aa9")
var MARD = common.HexToAddress("0x79bb4d71f6c168531a259dd6d40a8d5de5a34427")
var SYL_TEST = common.HexToAddress("0x050d0e98aeb9a133552409ed01cc1affa7cb3b99")
var SYL = common.HexToAddress("0xdb8cb9905d81eebc77bac749db9a0b8ce6ab9bef")
var NolzeBoost = common.HexToAddress("0xb83d3643837f343139575abe0bed708b1050eec8")
var MetaBusBoost = common.HexToAddress("0x19e8364fd51895eb9e0a9f3cfecc676326705588")
var KTX = common.HexToAddress("0x19fee5357d737db1db88c407ef1fe5207cd77df6")
var wsKRNO = common.HexToAddress("0xe944134903694ebdbb56aadcfbdf400fb52ea487")
var KBI = common.HexToAddress("0xfe16c99551fb4125cb6162e81048c0650db44266")
var STF = common.HexToAddress("0xb236f3baffc8c9253ae7e5fb58ae5e0689f05e06")
var LBK = common.HexToAddress("0xdcd9c56af7c05194d3a8c4187262130759e91320")
var RHEA = common.HexToAddress("0x0758fb651282581f86316514e8f5021493e9ed83")
var sKRNO = common.HexToAddress("0x6555f93f608980526b5ca79b3be2d4edadb5c562")
var TRT = common.HexToAddress("0x19cc7ae09d41273c0d6568395454d9b27da21d15")
var MBX = common.HexToAddress("0xd068c52d81f4409b9502da926ace3301cc41f623")
var KSTAR = common.HexToAddress("0x07ffbdba745f3a98ec462385aedcdcd973021671")
var MTF = common.HexToAddress("0x8755b9b884fd369dcf1639239c2b44ca4e9a2d78")
var SPRX = common.HexToAddress("0xa749a5fc259a95e81d8579ed37ab5fe14aa39c35")
var CX2 = common.HexToAddress("0x4440e2aa546e9faeb564ff7f45a845aad06de651")
var HALF = common.HexToAddress("0xd6aa0b0f8ebba0c14661fea0a5e12a50cf349251")
var GRD = common.HexToAddress("0x0facf2288dd04707c8c9ae2353a5d92a220a0812")
var GAIA = common.HexToAddress("0xe2541f0c54202fcdad60523fab8bfaa2d2540115")
var GATO = common.HexToAddress("0xc4edcfed08a169342b479578b01c77efe32630ec")
var MITX = common.HexToAddress("0x448829e85c430477a715400d4a61aa808d3c900f")
var ORB = common.HexToAddress("0x01ad62e0ff6dcaa72889fca155c7036c78ca1783")
var LAY = common.HexToAddress("0x1223baf4f5fb9c9002a2154262440b9ed09d01a7")
var GHUB = common.HexToAddress("0x4836cc1f355bb2a61c210eaa0cd3f729160cd95e")
var WRS = common.HexToAddress("0xfa86afa48e9306826010bc11977cfdb827c727dd")
var DAI = common.HexToAddress("0x5fac6bcdc22e40906346c5a64be3a6dbec9207fa")
var MONETA = common.HexToAddress("0x5afda70db64de4d5d24e4e87a40bc5f429736bc5")
var TINDER = common.HexToAddress("0x7cd4a64946e91989a21548362c18052704fe5ed6")
var HSC = common.HexToAddress("0xcc8088361c6f6ea58ce2d121f84afb2115d12cc5")
var MSHARE = common.HexToAddress("0xfd7e99ab9822edc9b92e1f89a8a0b96a76e5740e")
var oMATIC = common.HexToAddress("0xa006ba407cfc6584c90bac24ed971261885a0fd6")
var Pome = common.HexToAddress("0x9a943f3f84afa673d2a7cf053b8192e372600f57")
var AOA = common.HexToAddress("0x8a14d0bde789e924ad255a82041c7f761d1c0029")
var oDKA = common.HexToAddress("0xd01d650a5920fc714b2f8ed9d53e3ffc663302e9")
var oMESH = common.HexToAddress("0x127a75b751ba810e459121af6207d83841c586b7")
var USDF = common.HexToAddress("0x92fdeca1a5f6a6a88ea3fd732cfc24b4c14befb5")
var FLEX = common.HexToAddress("0x06a3172fea7edf5cfe2da1bcefe1ffe4a917c059")
var oPOLA = common.HexToAddress("0x2bf4d89453deb5781fccc656a4cac60710af766d")
var MD = common.HexToAddress("0x9d3ae9a8449e2474df6a38d4c9f1a15c733702a0")
var ZTC = common.HexToAddress("0xd6243f133ebf7ea191fb0eb47017b809b46b15f1")
var HZF = common.HexToAddress("0x8beb4953d25838e44abafdc7f7634f930588198d")
var KONE = common.HexToAddress("0xcd65df1a138f0a4e989336ca64b0498f5d26777c")
var kpEKL = common.HexToAddress("0x08644836b786b69a5082fd4644a3f2d1534b11a8")
var KPG = common.HexToAddress("0xf05d180a169418959a017865866f0abaf7db7ead")
var KCLNK = common.HexToAddress("0xfbe0862f1591b2197bbb40970d99b8d20a57f8d5")
var TOX = common.HexToAddress("0x604d6a7c492b4953bd20e007c7220d6c3d867fc5")
var oXDT = common.HexToAddress("0xaeeca95c899660dc74886168d0ffdebf3669179d")
var ACA = common.HexToAddress("0xc925f8da1b430334d15fbeba8896c98a098bafc3")
var nice = common.HexToAddress("0x23b40267e0c526fe4279e507e78d4aa7760d53e6")
var KIRI = common.HexToAddress("0xeaca2a817632c0ab2ac608711e12ac5209d56da8")
var ABC = common.HexToAddress("0xce899f5fcf55b0c1d7478910f812cfe68c5bcf0f")
var ENRG = common.HexToAddress("0xe91ffe2e15ccd56b1b8ddf7cdf848dfee6b5a858")
var RSTAR = common.HexToAddress("0xe5f59ea8b7c9806dc84e8f0862e0f6176f2f9cf2")
var JAMBA = common.HexToAddress("0xaf89ca6f4ee85ff31dbe4df6bc50a2779f6ba820")
var DRB = common.HexToAddress("0x945f68b51cc51709f771e7104990b3f8a3c3ec79")
var MINT = common.HexToAddress("0xf092acc2412742f4d5a457799dea57155ed42f9c")
var SIG = common.HexToAddress("0x94a2a6308c0a3782d83ad590d82ff0ffcc515312")
var MSP = common.HexToAddress("0x922584e0f30883c467d6aed91dc65697664a7878")
var MEAT = common.HexToAddress("0x9f60e3b785b698758542f9a137141fdcb52eb6b2")
var TEST2848 = common.HexToAddress("0x2fb10a633f7dcd1970c4f88eaf8d8fddd4b20885")
var CITY = common.HexToAddress("0x52f9ad2033c6375bb3365494b9be5561e8161963")
var RST = common.HexToAddress("0x906df72b0f2f342e645479b7cb6693c6d78ef868")
var CNT = common.HexToAddress("0x5ab03cdb98ec84846a418d4c7cb1d481a1ef5818")
var NoA = common.HexToAddress("0x5d344cb62452fb1c01587ff85af9ffe7179397dc")
var ICT = common.HexToAddress("0x2b5065d6049099295c68f5fcb97b8b0d3c354df7")
var PLWI = common.HexToAddress("0x6bae4b6afc2856b4ac0fb1165cf85c4923302ba2")
var sigKSP = common.HexToAddress("0x6026c432c420dce0e7bc5f84b9df1637b9ce953b")
var SIDO = common.HexToAddress("0xc3130105c6e1df04116f4afec754e3891c69b30f")
var MON = common.HexToAddress("0xa8598d1d1e6e5ecf03fc236df3561d276038c174")
var MAG = common.HexToAddress("0x3a4019bbbca9e4dd4e5a2c522e77619f79497a05")
var KLE = common.HexToAddress("0x4021f0a5762ef5a1be8e57f0d4c684a9637744d2")
var BBT = common.HexToAddress("0x450afb14016c338145ece2d6080cf1ae563d8cee")
var SBT = common.HexToAddress("0x54736bf96b0d8e87cc1893b07a06b8ae96380223")
var TAT = common.HexToAddress("0x3a9aaa06e25d56610f3e4685d413bedc3bb77731")
var TBT = common.HexToAddress("0xed0aec25f955f2ab1fbaa684d8c77d71ab456b0f")
var Hodoo = common.HexToAddress("0x868db0c72fd98d1e069008a9d84ac4df671588c3")
var ZEMIT = common.HexToAddress("0x9d52704cd67d586ed2870d810b0cef2cc168ae42")
var MOOI = common.HexToAddress("0x4b734a4d5bf19d89456ab975dfb75f02762dda1d")
var WALK = common.HexToAddress("0x976232eb7eb92287ff06c5d145bd0d1c033eca58")
var MT = common.HexToAddress("0x593e493421ac12685b2dfccff27b62a2718bfec8")
