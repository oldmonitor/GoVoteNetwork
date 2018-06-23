package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type Wallet struct {
	Balance   float64
	keyPair   RSAKeyPair
	PublicKey string
}

func (w *Wallet) initWallet() {
	//for testing purpose, give everyone 100 initial balance
	w.Balance = 100

}

func (w *Wallet) generateKeyForWallet() {
	var keyPair RSAKeyPair
	keyPair.generateKey()

	w.keyPair = keyPair
	w.PublicKey = hex.EncodeToString(keyPair.PublicKey)
}

func (w Wallet) toString() string {
	output := fmt.Sprintf(`Wallet-
		Balance: %s
		Public Key: %s`,
		string(strconv.FormatFloat(w.Balance, 'E', -1, 64)),
		w.PublicKey)
	return output
}
