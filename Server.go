package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var bc Blockchain

//Server for handling web request
type Server struct {
	httpPort string
}

func (s Server) run() error {
	mux := s.makeMuxRouter()

	bc.addBlock([]byte("new transaction 1"))
	log.Println("Listening on ", s.httpPort)
	ser := &http.Server{
		Addr:           ":" + s.httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := ser.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s Server) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", s.handleGetBlockchain).Methods("GET")
	//muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func (s Server) handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
