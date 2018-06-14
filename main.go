package main

import (
	"fmt"
	"os"
	"strconv"
)

//command - GoChain p2pPortNumber
func main() {
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var s Server
	s.blockChain.initBlockChain()
	s.httpPort = os.Getenv("HTTP_PORT")
	s.run()
	*/
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

	var p P2pServer
	p.Initialize("DataFiles\\peers.txt", portNumber)
	p.StartServer()
}
