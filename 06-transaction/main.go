package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	//get nonce for a1 from system
	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	amout := big.NewInt(10000000000000)
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, a2, amout, 21000, gasprice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2022-06-12T14-25-14.179365454Z--834bec369b146648ec0e4170705d3a2725ab71be")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
