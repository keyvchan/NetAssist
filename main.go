package main

import (
	"fmt"
	"log"
	"os"

	"github.com/keyvchan/NetAssist/client"
	"github.com/keyvchan/NetAssist/server"
)

func main() {
	args := os.Args
	fmt.Println(args)
	types := args[1]
	switch types {
	case "server":
		server.Serve(args[2:])
	case "client":
		client.Req(args[2:])
	default:
		log.Fatal("Unknown type: ", types)
	}
}
