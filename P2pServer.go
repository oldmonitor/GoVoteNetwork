package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

//P2PServer - struct for server
type P2PServer struct {
	P2pPort    int //p2p port. It is being use for p2p communication
	Peers      []P2PPeer
	upgrader   websocket.Upgrader
	blockChain *Blockchain
}

//P2PPeer - struct for connect peer node
type P2PPeer struct {
	PeerAddress         string
	HTTPPort            int
	WebSocketConnection *websocket.Conn
	IsConnected         bool
}

//StartServer listen on given port. Initialize must be called first.
func (s *P2PServer) startServer() {

	//upgrader with buffer size
	s.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	dialer := websocket.Dialer{}

	//for each peer address, dial the address of each client.
	for i := 0; i < len(s.Peers); i++ {
		println("Try connecting to peer " + s.Peers[i].PeerAddress + "/ws")
		conn, _, err := dialer.Dial(s.Peers[i].PeerAddress+"/ws", nil)

		//if there is an error connecting to client, output error, else listen for message from connected peer
		if err != nil {
			s.removeUnresponsivePeer(s.Peers[i].PeerAddress)
			i--
			println(err.Error())
		} else {
			s.Peers[i].WebSocketConnection = conn
			s.Peers[i].IsConnected = true

			//listen for message from client
			go s.wsListen(s.Peers[i])
		}
	}

	//start listening on given port
	http.HandleFunc("/ws", s.wsHandler)
	println("P2P Listening on port " + strconv.Itoa(s.P2pPort))
	err := http.ListenAndServe(":"+strconv.Itoa(s.P2pPort), nil)
	for err != nil {
		println(err.Error())
	}

	serverWaitGroup.Done()
}

//wsListen - start listening for message on open connections
func (s *P2PServer) wsListen(peer P2PPeer) {

	println(peer.WebSocketConnection.RemoteAddr().String() + " Connected")
	println("Peers#: " + strconv.Itoa(s.getConnectedPeerCount()))
	for {
		//get the message
		messageType, p, err := peer.WebSocketConnection.ReadMessage()

		if err != nil {
			//if there is connection error, remove peer and display message and stop listening for message
			println(peer.WebSocketConnection.RemoteAddr().String() + " Disconnected")
			s.removeDisconnectedPeer(peer.WebSocketConnection)
			println("Peers: " + strconv.Itoa(s.getConnectedPeerCount()))
			return
		}

		//parse blockchain message
		var bc Blockchain
		err = json.Unmarshal([]byte(p), &bc)

		//process blockchain
		println("Message received: type " + strconv.Itoa(messageType) + ":" + string(p))
		for i := 0; i < len(bc.Chain); i++ {
			println(bc.Chain[i].toString())
		}

		//replace blockchain
		s.blockChain.replaceChain(bc)
	}
}

//removePeer removed connection from connectedPeers array
func (s *P2PServer) removeDisconnectedPeer(conn *websocket.Conn) bool {
	for i, v := range s.Peers {
		if v.WebSocketConnection == conn {
			s.Peers = append(s.Peers[:i], s.Peers[i+1:]...)
			return true
		}
	}
	return false
}

//removeUnresponsivePeer remove unresponseive peer from the collection
func (s *P2PServer) removeUnresponsivePeer(address string) bool {
	for i, v := range s.Peers {
		if v.PeerAddress == address {
			s.Peers = append(s.Peers[:i], s.Peers[i+1:]...)
			return true
		}
	}
	return false
}

//get total number of connected peer
func (s P2PServer) getConnectedPeerCount() int {
	var count int
	for _, v := range s.Peers {
		if v.IsConnected == true {
			count++
		}
	}
	return count
}

//wsHanlder - web socket handler. This method is triggered when connection is establish.
//send current blockchain to peer once connection is established
func (s *P2PServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	s.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		println(err.Error())
		return
	}

	//add the connected peer to collection
	var peer P2PPeer
	peer.WebSocketConnection = conn
	peer.IsConnected = true
	s.Peers = append(s.Peers, peer)

	//send a blockchain to newly connected peer
	s.sendBlockchainToPeer(peer)

	//listen for message from peer
	s.wsListen(peer)
}

func (s *P2PServer) sendBlockchainToPeer(peer P2PPeer) {

	println("sending block chain to peer ", peer.PeerAddress)
	bc := s.blockChain
	peer.WebSocketConnection.WriteJSON(bc)
}

//syncBlockchain sends updated blockchain to all connected peer. This method is called when a new block
//is mined
func (s *P2PServer) syncBlockchain() {
	for i := 0; i < len(s.Peers); i++ {
		if s.Peers[i].IsConnected {
			s.sendBlockchainToPeer(s.Peers[i])
		}
	}
}
