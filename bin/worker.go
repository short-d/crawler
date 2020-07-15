package main

import (
	"crawler/worker"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Usage: worker worker_port master_ip master_port")
		return
	}

	workerPort, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	masterIP := args[1]
	masterPort, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	server := worker.NewServer(masterIP, masterPort, "encrypted")
	err = server.Start(workerPort)
	if err != nil {
		panic(err)
	}
}
