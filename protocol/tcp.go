package protocol

import (
	"net"
	"os"
	"strconv"

	"github.com/keyvchan/NetAssist/pkg/connection"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
	"github.com/rs/zerolog/log"
)

// TCPServer is a TCP server, read from stdin and write to the client and read from the client write it to stdout
func TCPServer() {
	address := flags.Config.Host + ":" + strconv.Itoa(flags.Config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Err(err).Msg("Could not listen on address")
	}
	// store all connections in a slice
	// NOTE: Possibly race condition
	var connections = map[net.Conn]bool{}

	// create a chennel to communicate between the read and write goroutines
	conn_read := make(chan message.Message)
	stdin_read := make(chan message.Message)

	stdin := connection.File{
		Path: "/dev/stdin",
		File: os.Stdin,
	}

	// create a goroutine to read user input
	go message.Read(stdin_read, stdin)
	go message.Write(conn_read, connection.Stdout)

	// create a goroutines cleanup closed conn
	go conn_cleanup(*connection.ClosedConn, connections)
	go write_to_conns(stdin_read, connections)
	go accept_conn(conn_read, listener, connections)
	select {}

}

// accept_conn accepts connections from the listener and adds them to the connections map
func accept_conn(read_chan chan message.Message, listener net.Listener, connections map[net.Conn]bool) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Err(err).Msg("Could not accept connection")
		}
		log.Info().Msg("Accepted connection")
		connections[conn] = true
		// create conn
		tcp_client := connection.Stream{
			Type: "tcp",
			Conn: conn,
		}
		go message.Read(read_chan, tcp_client)
	}

}

// TCPClient is a TCP client, read from stdin and write to the server and read from the server when it to stdout
func TCPClient() {
	address := flags.Config.Host + ":" + strconv.Itoa(flags.Config.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Err(err).Msg("Could not connect to server")
	}
	defer conn.Close()

	tcp_client := connection.Stream{
		Type: "tcp",
		Conn: conn,
	}

	// create a chennel to communicate between the read and write goroutines
	stdin_read := make(chan message.Message)
	go message.Read(stdin_read, connection.Stdin)
	go message.Write(stdin_read, tcp_client)

	quit := make(chan bool)
	conn_read := make(chan message.Message)

	go message.Read(conn_read, tcp_client)
	go message.Write(conn_read, connection.Stdout)
	go func() {
		conn := <-*connection.ClosedConn
		conn.Close()
		quit <- true
	}()

	// quit on signal
	<-quit
}

// conn_cleanup removes closed connections from the connections map
func conn_cleanup(closed_conn chan net.Conn, conns map[net.Conn]bool) {
	for {
		conn := <-closed_conn
		conn.Close()
		delete(conns, conn)
	}

}

// write_to_conns writes messages to all connections in the connections map
func write_to_conns(message_chan chan message.Message, connections map[net.Conn]bool) {
	for {
		message := <-message_chan
		// check message
		for conn := range connections {
			conn.Write(message.Content)
		}
	}
}
