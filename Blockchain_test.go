package main

import (
	"strings"
	"testing"
	"time"
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

	if block1.Hash != block2.Lasthash {
		t.Errorf("Expect the previous has value %v, but got %v", block1.Hash, block2.Lasthash)
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

	if chain[1].Hash == chain[2].Lasthash && bc.validateChain() == false {
		t.Errorf("Expect the hash value of first two blocks to match")
		return
	}

	//manually change hash value, the validation shoudl fail
	chain[1].Hash = chain[1].Hash + "111"
	if chain[1].Hash != chain[2].Lasthash && bc.validateChain() == true {
		t.Errorf("Expect the hash value of first two blocks not to match and validation fail")
		return
	}

}

//longer chain should replace current chain
func TestTryingToReplaceWithLongerChain(t *testing.T) {
	var bc Blockchain
	bc.addBlock([]byte("new transaction 1"))
	bc.addBlock([]byte("new transaction 1"))
	chain := bc.Chain

	var bcLongerChain Blockchain
	bcLongerChain.Chain = chain

	bcLongerChain.addBlock([]byte("new transaction 3"))

	var beforeLength = len(bc.Chain)
	bc.replaceChain(bcLongerChain)
	var afterLength = len(bc.Chain)

	if afterLength <= beforeLength {
		t.Errorf("Chain did not get replaced by longer chain")
		return
	}
}

//shorter chain should not replace current chain
func TestTryingToReplaceWithShorterChain(t *testing.T) {
	var bc Blockchain
	bc.addBlock([]byte("new transaction 1"))
	bc.addBlock([]byte("new transaction 1"))
	chain := bc.Chain

	var bcShorterChain Blockchain
	bcShorterChain.Chain = chain

	bc.addBlock([]byte("new transaction 3"))

	var beforeLength = len(bc.Chain)
	bc.replaceChain(bcShorterChain)
	var afterLength = len(bc.Chain)

	if afterLength != beforeLength {
		t.Errorf("Chain get replaced by shorter chain")
		return
	}
}

//if mine with diffulty level 2, the hash value should start with two 0
func TestMiningWithDiffultyLevelTwo(t *testing.T) {
	var bc Blockchain
	bc.initBlockChain()
	bc.Chain[0].Difficulty = 2
	bc.addBlock([]byte("new transaction 1"))
	var chainLength = len(bc.Chain)
	var lastBlock = bc.Chain[chainLength-1]

	if strings.HasPrefix(lastBlock.Hash, strings.Repeat("0", lastBlock.Difficulty)) == false {
		t.Errorf("The hash of last block is not valid")
		return
	}

	return
}

//if the block
func TestDynamicAdjustDiffultyLevel(t *testing.T) {
	var newBlock = Block{
		Lasthash:   "0000000",
		Data:       []byte("This is test"),
		Nonce:      0,
		Difficulty: 2,
	}

	var mRate int
	mRate = 3000
	originalDifficulty := newBlock.Difficulty
	newBlock.Timestamp = time.Now().Add(time.Millisecond * time.Duration(mRate*-2))
	adjustDifficulty(&newBlock, mRate)
	if newBlock.Difficulty <= originalDifficulty {
		t.Errorf("Expect the difficulty level increase to 3")
	}

	newBlock.Difficulty = 2
	newBlock.Timestamp = time.Now().Add(time.Millisecond * time.Duration(mRate/-2))
	adjustDifficulty(&newBlock, mRate)
	if newBlock.Difficulty >= originalDifficulty {
		t.Errorf("Expect the difficulty level decrease to 1")
	}

}
