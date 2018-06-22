package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var serverWaitGroup sync.WaitGroup

type BlockchainClient struct {
	p2pServer       P2PServer
	httpServer      HTTPServer
	blockchain      Blockchain
	isNewBlockAdded bool
}

func (bc *BlockchainClient) initClient(peersListFilePath string, p2pPortNumber int, httpPortNumber int) {
	//initialize the p2pserver
	bc.initP2pServer(peersListFilePath, p2pPortNumber)

	//initialize the httpserver
	bc.initHTTPServer(httpPortNumber)

	bc.initBlockchain()

}

func (bc *BlockchainClient) start() {
	serverWaitGroup.Add(2)

	//start p2p server
	go bc.p2pServer.startServer()

	//start http server
	go bc.httpServer.startServer()

	//listen for updated chain
	go bc.listenNewBlockMined()

	serverWaitGroup.Wait()
}

func (bc *BlockchainClient) initBlockchain() {
	bc.blockchain.initBlockChain()
}

func (bc *BlockchainClient) initHTTPServer(httpPortNumber int) {
	bc.httpServer.httpPort = httpPortNumber
	bc.httpServer.blockChain = &bc.blockchain
	bc.httpServer.isNewBlockAdded = &bc.isNewBlockAdded
}

func (bc *BlockchainClient) initP2pServer(peersListFilePath string, p2pPortNumber int) {
	file, err := os.Open(peersListFilePath)
	bc.p2pServer.P2pPort = p2pPortNumber
	bc.p2pServer.blockChain = &bc.blockchain

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

//listen for flag indicating a new block has been mined
func (bc *BlockchainClient) listenNewBlockMined() {
	println("Start checking if new block is mined")
	for {
		time.Sleep(10 * time.Millisecond)
		if bc.isNewBlockAdded {
			println("New block is mined. Sync chain to connected clients")

			//send the updated chain to all connected peers
			bc.p2pServer.syncBlockchain()
			bc.isNewBlockAdded = false
		}
	}

}
