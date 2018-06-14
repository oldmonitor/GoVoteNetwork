package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var s Server
	s.blockChain.initBlockChain()
	s.httpPort = os.Getenv("HTTP_PORT")
	s.run()

}
