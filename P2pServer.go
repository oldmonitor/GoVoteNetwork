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
	P2pPort        int
	PeerAddresses  []string
	ConnectedPeers []*websocket.Conn
	upgrader       websocket.Upgrader
}

//Initialize the p2p client/server
func (s *P2pServer) Initialize(peersConfigFileName string, p2pPort int) {

	//init peers config setting
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
		//ws://localhost:5002
		if scanner.Text() == "ws://localhost:"+strconv.Itoa(s.P2pPort) {
			continue
		}
		s.PeerAddresses = append(s.PeerAddresses, scanner.Text())
		fmt.Println(scanner.Text())
	}

	//upgrader with buffer size
	s.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

//StartServer listen on given port. Initialize must be called first.
func (s *P2pServer) StartServer() {

	//dial the known peers
	dialer := websocket.Dialer{}
	for _, arr := range s.PeerAddresses {
		println("try connecting " + arr + "/ws")
		conn, _, err := dialer.Dial(arr+"/ws", nil)

		//if there is an error connecting to client, output error, else listen for message from connected peer
		if err != nil {
			println(err.Error())
		} else {
			go s.wsListen(conn)
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

//wsListen - start listening on open connection
func (s *P2pServer) wsListen(conn *websocket.Conn) {

	//add connection to ConnectedPeers array
	s.ConnectedPeers = append(s.ConnectedPeers, conn)
	println(conn.RemoteAddr().String() + " Connected")
	println("Peers#: " + strconv.Itoa(len(s.ConnectedPeers)))
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			//if there is connection error, remove peer and display message
			println(conn.RemoteAddr().String() + " Disconnected")
			s.removePeer(conn)
			println("Peers: " + strconv.Itoa(len(s.ConnectedPeers)))
			return
		}
		//print message
		println(string(p))
		if err = conn.WriteMessage(messageType, p); err != nil {
			println(err.Error())
			return
		}
	}
}

//removePeer removed connection from connectedPeers array
func (s *P2pServer) removePeer(conn *websocket.Conn) bool {
	for i, v := range s.ConnectedPeers {
		if v == conn {
			s.ConnectedPeers = append(s.ConnectedPeers[:i], s.ConnectedPeers[i+1:]...)
			return true
		}
	}
	return false
}

//wsHanlder - web socket handler
func (s *P2pServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	s.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		println(err.Error())
		return
	}
	s.wsListen(conn)
}
