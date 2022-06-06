package protocol

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/pkg/utils"
)

func UDPServer() {
	address := internal.GetArg(3)
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	stdin_read_chan := make(chan internal.Message)
	conn_read_chan := make(chan internal.Message)

	go utils.ReadMessage(stdin_read_chan, os.Stdin)
	go utils.ReadMessage(conn_read_chan, conn)
	go utils.WriteMessage(conn_read_chan, os.Stdout)
	go utils.WriteMessage(stdin_read_chan, conn)
	select {}
}

func UDPClient() {
	address := internal.GetArg(3)
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stdin_read_chan := make(chan internal.Message)
	conn_read_chan := make(chan internal.Message)
	// we don't need to close udp socket

	go utils.ReadMessage(stdin_read_chan, os.Stdin)
	go utils.ReadMessage(conn_read_chan, conn)
	go utils.WriteMessage(conn_read_chan, os.Stdout)
	go utils.WriteMessage(stdin_read_chan, conn)

	select {}
}
