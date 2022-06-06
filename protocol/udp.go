package protocol

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/pkg/connection"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
	"github.com/keyvchan/NetAssist/pkg/utils"
)

func UDPServer() {
	address := flags.GetArg(3)
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	stdin_read_chan := make(chan message.Message)
	conn_read_chan := make(chan message.Message)

	udp_client_conn := connection.Packet{
		Type:       "udp",
		PacketConn: conn,
	}

	go utils.ReadMessage(stdin_read_chan, connection.Stdin)
	go utils.ReadMessage(conn_read_chan, udp_client_conn)
	go utils.WriteMessage(conn_read_chan, os.Stdout)
	go utils.WriteMessage(stdin_read_chan, conn)
	select {}
}

func UDPClient() {
	address := flags.GetArg(3)
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stdin_read_chan := make(chan message.Message)
	conn_read_chan := make(chan message.Message)
	// we don't need to close udp socket
	udp_server_conn := connection.Stream{
		Type: "udp",
		Conn: conn,
	}

	go utils.ReadMessage(stdin_read_chan, connection.Stdin)
	go utils.ReadMessage(conn_read_chan, udp_server_conn)
	go utils.WriteMessage(conn_read_chan, os.Stdout)
	go utils.WriteMessage(stdin_read_chan, conn)

	select {}
}
