package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/2ee533b24d4d4bbd97523ecf4113e8bf"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal("error creating")
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal("error creating block %v", err)
	}

	fmt.Println("the block number :", block.Number())

	addr := "0x093b47070640aa0388f36ae469a14b8cc65d137e"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		log.Fatal("error creating balance %v", err)
	}

	fmt.Println("The balance: ", balance)

	fBlance := new(big.Float)
	fBlance.SetString(balance.String())
	fmt.Println(fBlance)

	value := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println("The value: ", value)

}
