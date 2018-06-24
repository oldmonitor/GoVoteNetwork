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

func (w Wallet) sign(tranOutputs []TransactionOutput) string {
	hashValue := w.CreateHashOfTransactionOutput(tranOutputs)
	signedHash := rsaSign(w.keyPair.PrivateKey, []byte(hashValue))
	return hex.EncodeToString(signedHash)
}

func (w Wallet) verifySignature(t Transaction) bool {
	hashValue := w.CreateHashOfTransactionOutput(t.outputs)
	signature, _ := hex.DecodeString(t.input.signature)
	message := []byte(hashValue)

	isVerified := rsaUnsign([]byte(w.PublicKey), message, signature)

	return isVerified
}

func (w Wallet) CreateHashOfTransactionOutput(tranOutputs []TransactionOutput) string {
	var rawData []byte
	var newData []byte
	for _, element := range tranOutputs {
		newData = []byte(element.address + strconv.FormatFloat(element.amount, 'f', 2, 64))
		rawData = append(newData, rawData...)
	}
	hashValue := createHash(rawData)
	return hashValue
}

func (w Wallet) toString() string {
	output := fmt.Sprintf(`Wallet-
		Balance: %s
		Public Key: %s`,
		string(strconv.FormatFloat(w.Balance, 'E', -1, 64)),
		w.PublicKey)
	return output
}
