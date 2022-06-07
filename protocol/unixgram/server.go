package unixgram

import (
	"fmt"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/flags"
)

// Server is a server for the unixgram protocol, its create a bridge between server and client
func Server() {
	address := flags.GetArg(3)
	conn, err := net.ListenPacket("unixgram", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		conn.ReadFrom(buf)
		fmt.Println(string(buf))
	}
}
