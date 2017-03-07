package main

import (
	"fmt"
)

var nodeID string

var clip *CLIParser

var server *Server

var p2pManager *P2PManager

func main() {
	// parsing command line arguments
	parseCLI()

	// there are two basic modes: server (bank) and client (default)
	if clip.ServerMode {
		initServer()
		server.listen()
	} else {
		fmt.Println("Client - TODO")
	}

	fmt.Println(nodeID)
}
