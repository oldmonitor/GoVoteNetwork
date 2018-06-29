package main

import (
	"errors"
	"fmt"
	"time"
)

//Transaction struct
type Transaction struct {
	id      string //unique identifier of transaction
	input   TransactionInput
	outputs []TransactionOutput
}

//TransactionInput - addres from which money was sent
type TransactionInput struct {
	timestamp time.Time //timestamp
	amount    float64   //transfer amount
	address   string    //sender address
	signature string    //sender signature
}

//TransactionOutput - address to which money was sent
type TransactionOutput struct {
	amount  float64
	address string
}

func createNewTransaction(senderWallet Wallet, recipient string, amount float64) (Transaction, error) {
	var newTransaction Transaction

	if senderWallet.Balance < amount {
		fmt.Println("Not enought money in wallet")
		return newTransaction, errors.New("Not enought money in wallet")
	}

	//sender
	newTransaction.outputs = append(newTransaction.outputs, TransactionOutput{
		address: senderWallet.PublicKey, amount: senderWallet.Balance - amount,
	})

	//recipient
	newTransaction.outputs = append(newTransaction.outputs, TransactionOutput{
		address: recipient, amount: amount,
	})

	//sign transaction
	signTransaction(&newTransaction, senderWallet)

	return newTransaction, nil
}

//put digit signature to transaction
func signTransaction(transaction *Transaction, wallet Wallet) {
	var tranInput TransactionInput
	tranInput.timestamp = time.Now()
	tranInput.amount = wallet.Balance
	tranInput.address = wallet.PublicKey
	tranInput.signature = wallet.sign(transaction.outputs)
	transaction.input = tranInput
}
