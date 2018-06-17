package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Server for handling web request
type HTTPServer struct {
	httpPort   string
	blockChain Blockchain
}

func (s *HTTPServer) run() error {
	mux := s.makeMuxRouter()

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

func (s *HTTPServer) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", s.handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", s.handleWriteBlock).Methods("POST")
	return muxRouter
}

func (s *HTTPServer) handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(s.blockChain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func (s *HTTPServer) handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Blockchain

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, m)
		return
	}
	defer r.Body.Close()

	s.blockChain.replaceChain(m)
	s.blockChain.addBlock([]byte("new transaction 1"))
	respondWithJSON(w, r, http.StatusCreated, s.blockChain)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload Blockchain) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
