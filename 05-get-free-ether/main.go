package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://kovan.infura.io/v3/2ee533b24d4d4bbd97523ecf4113e8bf"

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// "834bec369b146648ec0e4170705d3a2725ab71be"
	// "1de6640af50868465b7a9ad2a9cbbf28b1af8dd2"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	a1 := common.HexToAddress("834bec369b146648ec0e4170705d3a2725ab71be")
	a2 := common.HexToAddress("1de6640af50868465b7a9ad2a9cbbf28b1af8dd2")

	//get balance of address1 and address2
	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
}
