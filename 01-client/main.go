package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/2ee533b24d4d4bbd97523ecf4113e8bf"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURL)

	if err != nil {
		log.Fatal("error creating")
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal("error creating block %v", err)
	}

	fmt.Println(block.Number())

}
