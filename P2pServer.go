package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

//P2pServer - struct for server
type P2pServer struct {
	P2pPort int //p2p port. It is being use for p2p communication
	//PeerAddresses  []string
	//ConnectedPeers []*websocket.Conn
	Peers    []P2pPeer
	upgrader websocket.Upgrader
}

//P2pPeer - struct for connect peer node
type P2pPeer struct {
	PeerAddress         string
	HTTPPort            int
	WebSocketConnection *websocket.Conn
	IsConnected         bool
}

//StartServer listen on given port. Initialize must be called first.
func (s *P2pServer) StartServer(peersConfigFileName string, p2pPort int) {

	//initialize server config
	s.initialize(peersConfigFileName, p2pPort)

	dialer := websocket.Dialer{}

	//for each peer address, dial the known port
	for i := 0; i < len(s.Peers); i++ {
		println("try connecting " + s.Peers[i].PeerAddress + "/ws")
		conn, _, err := dialer.Dial(s.Peers[i].PeerAddress+"/ws", nil)

		//if there is an error connecting to client, output error, else listen for message from connected peer
		if err != nil {
			s.removeUnresponsivePeer(s.Peers[i].PeerAddress)
			i--
			println(err.Error())
		} else {
			s.Peers[i].WebSocketConnection = conn
			s.Peers[i].IsConnected = true
			go s.wsListen(s.Peers[i])
		}
	}

	//start listening on given port
	http.HandleFunc("/ws", s.wsHandler)
	println("Listening on port " + strconv.Itoa(s.P2pPort))
	err := http.ListenAndServe(":"+strconv.Itoa(s.P2pPort), nil)
	for err != nil {
		println(err.Error())
	}
}

//Initialize - initialize the p2p client/server
func (s *P2pServer) initialize(peersConfigFileName string, p2pPort int) {

	file, err := os.Open(peersConfigFileName)
	s.P2pPort = p2pPort

	defer file.Close()
	if err != nil {
		println(err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("Peers in data file:")
	for scanner.Scan() {
		//ignore local host with same port

		var peer P2pPeer
		peer.PeerAddress = scanner.Text()
		peer.IsConnected = false
		fmt.Println(scanner.Text())
		//if address is current local host, skip
		if peer.PeerAddress == "ws://localhost:"+strconv.Itoa(s.P2pPort) {
			continue
		}

		s.Peers = append(s.Peers, peer)

	}

	//upgrader with buffer size
	s.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

//wsListen - start listening on open connection
func (s *P2pServer) wsListen(peer P2pPeer) {

	println(peer.WebSocketConnection.RemoteAddr().String() + " Connected")
	println("Peers#: " + strconv.Itoa(s.getConnectedPeerCount()))
	for {
		messageType, p, err := peer.WebSocketConnection.ReadMessage()

		if err != nil {
			//if there is connection error, remove peer and display message
			println(peer.WebSocketConnection.RemoteAddr().String() + " Disconnected")
			s.removeDisconnectedPeer(peer.WebSocketConnection)
			println("Peers: " + strconv.Itoa(s.getConnectedPeerCount()))
			return
		}
		//print message
		println(string(p))
		if err = peer.WebSocketConnection.WriteMessage(messageType, p); err != nil {
			println(err.Error())
			return
		}
	}
}

//removePeer removed connection from connectedPeers array
func (s *P2pServer) removeDisconnectedPeer(conn *websocket.Conn) bool {
	for i, v := range s.Peers {
		if v.WebSocketConnection == conn {
			s.Peers = append(s.Peers[:i], s.Peers[i+1:]...)
			return true
		}
	}
	return false
}

//removeUnresponsivePeer remove unresponseive peer from the collection
func (s *P2pServer) removeUnresponsivePeer(address string) bool {
	for i, v := range s.Peers {
		if v.PeerAddress == address {
			s.Peers = append(s.Peers[:i], s.Peers[i+1:]...)
			return true
		}
	}
	return false
}

func (s P2pServer) getConnectedPeerCount() int {
	var count int = 0
	for _, v := range s.Peers {
		if v.IsConnected == true {
			count++
		}
	}
	return count
}

//wsHanlder - web socket handler
func (s *P2pServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	s.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		println(err.Error())
		return
	}

	var peer P2pPeer
	peer.WebSocketConnection = conn
	peer.IsConnected = true
	s.Peers = append(s.Peers, peer)
	s.wsListen(peer)
}
