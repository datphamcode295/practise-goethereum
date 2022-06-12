package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"

	// a, err := key.NewAccount(password)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(a)

	b, err := ioutil.ReadFile("./wallet/UTC--2022-06-12T14-25-14.179365454Z--834bec369b146648ec0e4170705d3a2725ab71be")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private : ", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public :", hexutil.Encode(pData))

	fmt.Println(crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
