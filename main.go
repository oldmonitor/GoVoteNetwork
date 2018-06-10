package main

import (
	"fmt"
)

func main() {
	var bc Blockchain
	var data []byte
	data = []byte("1111111111")
	bc.AddBlock(data)

	var chain []Block
	chain = bc.Chain

	for i := 0; i < len(chain); i++ {
		fmt.Println(chain[i].ToString())
	}

}
