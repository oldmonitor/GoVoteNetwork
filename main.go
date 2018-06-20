package main

import (
	"fmt"
	"os"
	"strconv"
)

//command - GoChain p2pPortNumber
func main() {

	fmt.Println("Starting ...")
	var p2pPortNumber int
	var httpPortNumber int

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

	//portNumber = 5002
	var client BlockchainClient
	var peerListFilePath = "DataFiles\\peers.txt"
	client.initClient(peerListFilePath, p2pPortNumber, httpPortNumber)
	client.start()

}
