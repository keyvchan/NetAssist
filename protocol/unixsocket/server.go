package unixsocket

import (
	"fmt"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/flags"
)

// UnixServer is a server for the unix socket, it bridged server to stdout and stdin to server
func Server() {
	address := flags.GetArg(3)
	listener, err := net.Listen("unix", address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	log.Println("Accepted connection")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	}

}
