package main

import "testing"

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
	if len(tra.output) != 2 {
		t.Errorf("there should be only 2 transaction outputs")
	}

	//for sender, the current balance - transfer amount should equal to the amount in transaction output
	for i := 0; i < len(tra.output); i++ {
		if tra.output[i].address == senderW.PublicKey {
			if tra.output[i].amount != (senderW.Balance - amount) {
				t.Errorf("sender amount not correct. Expecting %f. Received %f", (senderW.Balance - amount), tra.output[i].amount)
			}
		}
	}

	for i := 0; i < len(tra.output); i++ {
		if tra.output[i].address == recipient {
			if tra.output[i].amount != amount {
				t.Errorf("recipient amount not correct. Expecting %f. Received %f", amount, tra.output[i].amount)
			}
		}
	}

}
