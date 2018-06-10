package main

import (
	"time"
)

// Blockchain - chained transaction
type Blockchain struct {
	Chain []Block
}

// AddBlock add a block to current chain
func (bc *Blockchain) AddBlock(blockdata []byte) {
	//chain is empty, initialize chain with genesis block
	if bc.Chain == nil {
		genBlock := GetGenesisBlock()
		genBlock.EncryptData()
		bc.Chain = []Block{genBlock}
	}

	//create the new block. lastHash link the new block and last block together
	var lastBlock = bc.Chain[len(bc.Chain)-1]
	var newBlock = Block{
		timestamp: time.Now(),
		lasthash:  lastBlock.hash,
		data:      blockdata}

	newBlock.EncryptData()

	//append block to chain
	bc.Chain = append(bc.Chain, newBlock)
}

//ValidateChain check blockchain. if invalid, return false.
func (bc Blockchain) ValidateChain() bool {

	return true
}

// GetGenesisBlock Create genesis block. First block of the chain
func GetGenesisBlock() Block {
	var b Block
	b.timestamp = time.Now()
	b.hash = "0000000000"
	b.lasthash = "0000000000"
	b.data = []byte("0000000000")
	return b
}

// MineBlock Create a block for the chain
func MineBlock(lastBlock Block, data []byte) Block {
	var b Block
	b.timestamp = time.Now()
	b.hash = "0000000000"
	b.lasthash = lastBlock.hash
	b.data = data
	return b
}
