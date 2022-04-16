package main

import (
	"log"

	"github.com/keyvchan/NetAssist/client"
	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/server"
)

func main() {
	internal.SetArgs()
	types := internal.GetArg(1)
	switch types {
	case "server":
		server.Serve()
	case "client":
		client.Req()
	default:
		log.Fatal("Unknown type: ", types)
	}
}
