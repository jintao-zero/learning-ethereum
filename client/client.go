package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// local ganache
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	fmt.Println("get a eth client")
	bNum, err := client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("get latest block number", bNum)
	var blockNumber big.Int
	blockNumber.SetUint64(bNum)
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress("0x06ca0BeF13a26FFb03d64f23969548C1C958e7Fe"), &blockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of %s is %s Wei", "0x06ca0BeF13a26FFb03d64f23969548C1C958e7Fe", balance.String())
	fmt.Println()
}
