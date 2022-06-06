package protocol

import (
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/connection"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
)

func UDPServer() {
	address := flags.GetArg(3)
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	stdin_read_chan := make(chan message.Message)
	conn_read_chan := make(chan message.Message)

	udp_server_conn := connection.Packet{
		Type:       "udp",
		PacketConn: conn,
	}

	go message.Read(stdin_read_chan, connection.Stdin)
	go message.Read(conn_read_chan, udp_server_conn)
	go message.Write(conn_read_chan, connection.Stdout)
	go message.Write(stdin_read_chan, udp_server_conn)
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
	udp_client_conn := connection.Stream{
		Type: "udp",
		Conn: conn,
	}

	go message.Read(stdin_read_chan, connection.Stdin)
	go message.Read(conn_read_chan, udp_client_conn)
	go message.Write(conn_read_chan, connection.Stdout)
	go message.Write(stdin_read_chan, udp_client_conn)

	select {}
}
