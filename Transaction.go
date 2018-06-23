package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	id     string //unique identifier of transaction
	input  string
	output []TransactionOutput
}

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
	newTransaction.output = append(newTransaction.output, TransactionOutput{
		address: senderWallet.PublicKey, amount: senderWallet.Balance - amount,
	})

	//recipient
	newTransaction.output = append(newTransaction.output, TransactionOutput{
		address: recipient, amount: amount,
	})
	return newTransaction, nil
}
