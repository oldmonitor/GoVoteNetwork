package main

import (
	"fmt"
	"testing"
)

func TestWalletKey(t *testing.T) {
	var wallet Wallet
	wallet.initWallet()
	wallet.generateKeyForWallet()
	fmt.Println(wallet.PublicKey)
	if wallet.PublicKey == "" {
		t.Errorf("Wallet public key is null after initialization")
	}
}
