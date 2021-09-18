package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/984241efe32a44ef8ef9c8b94c9006f4")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0xe4d6bc24a155133EB83dC9485060E89ecF9AF86b")
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	/*
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
 	*/

	toAddress := common.HexToAddress("0x5d4C7803870a6AB2fC1816711751522a3171cEf8")
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	txData := types.DynamicFeeTx{
		ChainID: chainID,
		Nonce: nonce,
		GasTipCap: big.NewInt(10* params.GWei),
		GasFeeCap: big.NewInt(20 * params.GWei),
		Gas: 21000,
		To: &toAddress,
		Value: big.NewInt(88*params.GWei),
		Data: nil,
	}
	tx := types.NewTx(&txData)

	privateKey, err := crypto.HexToECDSA("8c42416f1d2a9ed15167fdc8faa6b40dd910813b471e26b2290e0eba79cb9e1f")
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal("sign", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("send ", err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
