package main

import (
	wallet2 "blockchain/wallet"
	"fmt"
)

func main() {

	wallet := wallet2.NewWallet()
	fmt.Println(wallet.PublicKeyStr())
	fmt.Println(wallet.PublicKeyStr())

}
