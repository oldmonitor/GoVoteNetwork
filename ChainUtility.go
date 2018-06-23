package main

import (
	"fmt"
	"os"
)

//general method to handle errors
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
