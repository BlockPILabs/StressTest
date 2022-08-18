package main

import (
	"Web3ContractTest/contract/erc"
	"Web3ContractTest/contract/factory"
	"Web3ContractTest/contract/pair"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var totalNum atomic.Uint64
var FailNum atomic.Uint64
var reqDelay atomic.Uint64
var myStartTime time.Time
var myEndTime time.Time
var factoryContract *factory.Storage

// getPrice Get the price of a specific contract trading pair
func getPrice(conn *ethclient.Client, coinA common.Address, coinB common.Address) {
	pollAddr, _ := factoryContract.TokenToPool(nil, coinA, coinB)
	pairContract, _ := pair.NewStorage(pollAddr, conn)
	coinAContract, _ := erc.NewStorage(coinA, conn)
	dec, _ := coinAContract.Decimals(nil)
	bigIntDecA := decimal.New(1, int32(dec)).BigInt()
	for {
		// calculating time
		startT := time.Now()
		_, err := pairContract.EstimatePos(nil, coinB, bigIntDecA)
		// Total number of add requests
		totalNum.Add(1)
		tc := time.Since(startT)
		reqDelay.Add(uint64(tc.Milliseconds()))

		if err != nil {
			FailNum.Add(1)
		}
		time.Sleep(1 * time.Second)
	}
}

func logOut() {
	fmt.Printf("start_time:     %v\n", myStartTime)
	fmt.Printf("end_time:       %v\n", myEndTime)
	fmt.Printf("running_time:   %v\n", myEndTime.Sub(myStartTime))
	fmt.Printf("total_requests: %d\n", totalNum.Load())
	seccessNum := totalNum.Load() - FailNum.Load()
	fmt.Printf("http_req_success_num :%d\n", seccessNum)
	u := float64(seccessNum) / float64(totalNum.Load()) * 100
	fmt.Printf("http_req_success_pre :%.2f%%\n", u)
	fmt.Printf("http_req_avg:     %.2f\n", float64(totalNum.Load())/vtime.Seconds())
	fmt.Printf("http_req_latency: %v\n", time.Duration(reqDelay.Load()/totalNum.Load())*time.Millisecond)
}

func init() {
	initFlag()
	initVer()
}

func main() {
	myStartTime = time.Now()
	conn, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}
	factoryContract, _ = factory.NewStorage(factoryAddr, conn)

	// print log
	go func() {
		time.Sleep(3 * time.Second)
		for {
			fmt.Printf("total requests :%d\n", totalNum.Load())
			time.Sleep(1 * time.Second)
		}
	}()

	swapTokenListsLen := len(swapTokenLists)
	selectNum := 0
	for i := 0; i < vuser; i++ {
		token := swapTokenLists[selectNum]
		go getPrice(conn, token.coinA, token.coinB)
		selectNum += 1
		if selectNum == swapTokenListsLen {
			selectNum = 0
		}
	}
	time.Sleep(vtime)

	myEndTime = time.Now()
	logOut()
}
