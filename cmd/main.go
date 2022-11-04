package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Block struct {
	Number string
}

func main() {
	client, err := ethclient.Dial("http://172.16.2.247:8545")
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}

	id, err := client.ChainID(context.Background())
	if err != nil {
		return
	}
	fmt.Println(id)

	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())

	account := common.HexToAddress("0x04c871abd6d9f237ba23a8257438e11715e3743c")
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
}
