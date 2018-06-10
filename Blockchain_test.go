package main

import (
	"testing"
)

//Test when adding new block to new chain, the len of chain should be 2 and the second item's previous hash value
//should be equal to the hash value of first item
func TestAddingBlockToNewChain(t *testing.T) {
	var bc Blockchain
	var data []byte
	data = []byte("new transaction 1")
	bc.addBlock(data)

	var chain = bc.Chain
	if len(chain) != 2 {
		t.Errorf("Expect new chain has length 2, but got %v", len(chain))
		return
	}

	var block1 = chain[0]
	var block2 = chain[1]

	if block1.hash != block2.lasthash {
		t.Errorf("Expect the previous has value %v, but got %v", block1.hash, block2.lasthash)
		return
	}
}

//test blockchain validation. The hash does not matched for two chained blocks, the validation must return false
func TestBlockChainValidation(t *testing.T) {

	//init chain
	var bc Blockchain
	bc.addBlock([]byte("new transaction 1"))
	bc.addBlock([]byte("new transaction 1"))

	chain := bc.Chain

	if chain[1].hash == chain[2].lasthash && bc.validateChain() == false {
		t.Errorf("Expect the hash value of first two blocks to match")
		return
	}

	//manually change hash value, the validation shoudl fail
	chain[1].hash = chain[1].hash + "111"
	if chain[1].hash != chain[2].lasthash && bc.validateChain() == true {
		t.Errorf("Expect the hash value of first two blocks not to match and validation fail")
		return
	}

}
