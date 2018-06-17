package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var serverWaitGroup sync.WaitGroup

type BlockchainClient struct {
	p2pServer  P2PServer
	httpServer HTTPServer
}

func (bc *BlockchainClient) initClient(peersListFilePath string, p2pPortNumber int, httpPortNumber int) {
	//initialize the p2pserver
	bc.initP2pServer(peersListFilePath, p2pPortNumber)

	//initialize the httpserver
	bc.initHttpServer(httpPortNumber)
}

func (bc *BlockchainClient) start() {
	serverWaitGroup.Add(2)

	//start p2p server
	go bc.p2pServer.startServer()

	//start http server
	go bc.httpServer.startServer()

	serverWaitGroup.Wait()
}

func (bc *BlockchainClient) initHttpServer(httpPortNumber int) {
	bc.httpServer.httpPort = httpPortNumber
}

func (bc *BlockchainClient) initP2pServer(peersListFilePath string, p2pPortNumber int) {
	file, err := os.Open(peersListFilePath)
	bc.p2pServer.P2pPort = p2pPortNumber

	defer file.Close()
	if err != nil {
		println(err.Error())
		return
	}

	//create a peer object for each entry in the config file
	scanner := bufio.NewScanner(file)
	fmt.Println("Peers in data file:")
	for scanner.Scan() {
		//ignore local host with same port
		var peer P2PPeer
		peer.PeerAddress = scanner.Text()
		peer.IsConnected = false
		fmt.Println(scanner.Text())
		//if address is current local host, skip
		if peer.PeerAddress == "ws://localhost:"+strconv.Itoa(p2pPortNumber) {
			continue
		}

		bc.p2pServer.Peers = append(bc.p2pServer.Peers, peer)

	}

}
