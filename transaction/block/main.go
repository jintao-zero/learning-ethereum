package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("latest header number", header.Number.String()) // 5671744

	blockNumber := big.NewInt(5)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block number", block.Number().Uint64())         // 5671744
	fmt.Println("block time", block.Time())                      // 1527211625
	fmt.Println("block difficulty", block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println("block hash", block.Hash().Hex())                // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("block tx number", len(block.Transactions()))    // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx number", count, "block hash", block.Hash().Hex()) // 144
}
