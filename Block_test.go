package main

import (
	"testing"
)

//Test when adding new block to new chain, the len of chain should be 2
func TestAddingBlockToChain(t *testing.T) {
	var bc Blockchain
	var data []byte
	data = []byte("1111111111")
	bc.AddBlock(data)

	var chain = bc.Chain
	if len(chain) != 2 {
		t.Errorf("Expect new chain has length 2, but got " + len(bc))
	}
}
