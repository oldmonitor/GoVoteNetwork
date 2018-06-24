# GoChain
Implemented in Golang, a blockchain P2P network for wallet and encrypted transaction

Prerequisites

Golang compiler: https://golang.org/dl/

Running

1. go build
2. Update web socket addresses in peers.txt. 
3. Start client. The first parameter is the P2P port. The second parameter is the HTTP port

  3.1. gochain 5001 8001

  3.2. gochain 5002 8002

4. Send a mine request to the first client. 
  http://localhost:8001/mine
  message body (json raw):
  {"Message":"[data for new block]"}
