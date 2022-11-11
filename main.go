package main

import (
	"log"

	"github.com/keyvchan/NetAssist/client"
	"github.com/keyvchan/NetAssist/perf"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/server"
)

func main() {
	flags.SetArgs()
	types := flags.GetArg(1)
	switch types {
	case "server":
		server.Serve()
	case "client":
		client.Req()
	case "perf":
		perf.ServerPerf()
	default:
		log.Fatal("Unknown type: ", types)
	}
}
