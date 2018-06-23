package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/asn1"
	"fmt"
	"strconv"
)

type Wallet struct {
	Balance   int
	keyPair   RSAKeyPair
	PublicKey []byte
}

func (w *Wallet) initWallet() {
	//for testing purpose, give everyone 100 initial balance
	w.Balance = 100

}

func (w *Wallet) generateKeyForWallet() {
	reader := rand.Reader
	bitSize := 2048

	//generate key pair
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	//save public key
	asn1Bytes, err := asn1.Marshal(key.PublicKey)
	w.PublicKey = asn1Bytes
	checkError(err)
}

func (w Wallet) toString() string {
	output := fmt.Sprintf(`Wallet-
		Balance: %s
		Public Key: %s`,
		string(strconv.Itoa(w.Balance)),
		w.PublicKey)
	return output
}
