package main

import (
	"crawler/master"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: master master_port")
		return
	}

	port, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	server := master.NewServer()
	err = server.Start(port)
	if err != nil {
		panic(err)
	}
}
