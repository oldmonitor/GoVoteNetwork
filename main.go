package main

import (
	"fmt"
	"os"
	"strconv"
)

var p2pPortNumber int
var httpPortNumber int
var peerListFilePath string

func init() {
	//p2pPortNumber = 3001
	//httpPortNumber = 8001
	fmt.Println("Starting ...")

	//check if os.Args are valid (note: first arg is the command itself.)
	if os.Args == nil || len(os.Args) <= 2 {
		fmt.Println("Need pass p2p port number and http port number. Blockchain [p2pPortNumber] [httpPortNumber]")
		return
	}

	if os.Args[1] != "" {
		p2pPortNumber, _ = strconv.Atoi(os.Args[1])
	}

	if os.Args[2] != "" {
		httpPortNumber, _ = strconv.Atoi(os.Args[2])
	}
	peerListFilePath = "DataFiles\\peers.txt"
}

//command - GoChain p2pPortNumber
func main() {

	var client BlockchainClient

	client.initClient(peerListFilePath, p2pPortNumber, httpPortNumber)
	client.start()
}
