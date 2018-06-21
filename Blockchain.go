package main

import (
	"time"
)

var default_difficulty int = 4

// Blockchain - chained transaction
type Blockchain struct {
	Chain []Block
}

func (bc *Blockchain) initBlockChain() {
	genBlock := getGenesisBlock()
	genBlock.encryptData()
	bc.Chain = []Block{genBlock}
}

// AddBlock add a block to current chain
func (bc *Blockchain) addBlock(blockdata []byte) {
	//chain is empty, initialize chain with genesis block
	if bc.Chain == nil {
		bc.initBlockChain()
	}

	//create the new block. lastHash link the new block and last block together
	var lastBlock = bc.Chain[len(bc.Chain)-1]

	var newBlock = mineBlock(lastBlock, blockdata)

	//append block to chain
	bc.Chain = append(bc.Chain, newBlock)
}

//ValidateChain check blockchain. if invalid, return false.
func (bc Blockchain) validateChain() bool {

	if len(bc.Chain) == 0 {
		return false
	}

	//if the first block is not genesis, return false
	if bc.Chain[0].Lasthash != "0000000000" && string(bc.Chain[0].Data) != string([]byte("0000000000")) {
		return false
	}

	//compare hash value of each chained blocks
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		lastBlock := bc.Chain[i-1]
		if currentBlock.Lasthash != lastBlock.Hash {
			return false
		}
	}

	return true
}

//replaceChain replace current chain with longer valid chain
func (bc *Blockchain) replaceChain(newBc Blockchain) {
	//if new blockchain is longer and
	if len(newBc.Chain) > len(bc.Chain) && newBc.validateChain() {
		bc.Chain = newBc.Chain
	}
}

// GetGenesisBlock Create genesis block. First block of the chain
func getGenesisBlock() Block {
	var b Block
	b.Timestamp = time.Now()
	b.Lasthash = "0000000000"
	b.Data = []byte("0000000000")
	b.encryptData()
	return b
}

// MineBlock Create a block for the chain
func mineBlock(lastBlock Block, blockData []byte) Block {
	var newBlock = Block{
		Timestamp: time.Now(),
		Lasthash:  lastBlock.Hash,
		Data:      blockData}

	newBlock.encryptData()

	return newBlock
}
