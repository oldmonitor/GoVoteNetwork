package main

import (
	"fmt"
	"os"
	"strconv"
)

//command - GoChain p2pPortNumber
func main() {

	fmt.Println("Starting ...")
	var portNumber int
	//check if os.Args has only one arg. (note: first arg is the command itself.)
	if os.Args == nil || len(os.Args) == 1 {
		fmt.Println("Need pass p2p port number.")
		return
	}
	if os.Args[1] != "" {
		portNumber, _ = strconv.Atoi(os.Args[1])
	}
	//portNumber = 5001

	var p P2pServer
	p.StartServer("DataFiles\\peers.txt", portNumber)

}
