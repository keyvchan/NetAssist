package protocol

import (
	"fmt"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/internal"
)

func UDPServer() {
	address := internal.GetArg(3)
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		conn.ReadFrom(buf)
		fmt.Println(string(buf))
	}
}

func UDPClient() {
	address := internal.GetArg(3)
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// we don't need to close udp socket
	go internal.ConnRead(conn, nil)

	for {
		buf := internal.StdinRead()
		_, err := conn.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
	}

}
