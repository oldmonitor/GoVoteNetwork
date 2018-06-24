package main

import (
	"testing"
)

//for sender, the current balance - transfer amount should equal to the amount in transaction output
//for recepient, the transfer amount should be in the transaction output
func TestSubtractAmountTransaction(t *testing.T) {
	var senderW Wallet
	senderW.initWallet()
	senderW.generateKeyForWallet()
	recipient := "recipient123"
	var amount float64 = 10
	var tra Transaction
	tra, _ = createNewTransaction(senderW, recipient, amount)

	//there should be only 2 transaction output
	if len(tra.outputs) != 2 {
		t.Errorf("there should be only 2 transaction outputs")
	}

	//for sender, the current balance - transfer amount should equal to the amount in transaction output
	for i := 0; i < len(tra.outputs); i++ {
		if tra.outputs[i].address == senderW.PublicKey {
			if tra.outputs[i].amount != (senderW.Balance - amount) {
				t.Errorf("sender amount not correct. Expecting %f. Received %f", (senderW.Balance - amount), tra.outputs[i].amount)
			}
		}
	}

	for i := 0; i < len(tra.outputs); i++ {
		if tra.outputs[i].address == recipient {
			if tra.outputs[i].amount != amount {
				t.Errorf("recipient amount not correct. Expecting %f. Received %f", amount, tra.outputs[i].amount)
			}
		}
	}
}

//application should not create transaction ouput if transaction amount is larger then current balance
func TestTransactionAmountLargerThanCurrentBalance(t *testing.T) {
	var senderW Wallet
	senderW.initWallet()
	senderW.generateKeyForWallet()
	recipient := "recipient123"
	var amount float64 = 999999
	var tra Transaction
	tra, _ = createNewTransaction(senderW, recipient, amount)

	if len(tra.outputs) > 0 {
		t.Errorf("transaction amount is larger than current balance, transaction output should be 0")
	}
}

func TestIfNewTransactionIsSigned(t *testing.T) {
	var senderW Wallet
	senderW.initWallet()
	senderW.generateKeyForWallet()
	recipient := "recipient123"
	var amount float64 = 50
	var tra Transaction
	tra, _ = createNewTransaction(senderW, recipient, amount)

	if tra.input.signature == "" {
		t.Errorf("signature in input transaction is empty")
		return
	}
}

func TestOriginalTransactionShouldPassSignatureValidation(t *testing.T) {
	var senderW Wallet
	senderW.initWallet()
	senderW.generateKeyForWallet()
	recipient := "recipient123"
	var amount float64 = 50
	var tra Transaction
	tra, _ = createNewTransaction(senderW, recipient, amount)
	if senderW.verifySignature(tra) == false {
		t.Errorf("verification of original transaction failed")
	}

}
