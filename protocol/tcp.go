package protocol

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/pkg/connection"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
)

func TCPServer() {
	address := flags.GetArg(3)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
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

func accept_conn(read_chan chan message.Message, listener net.Listener, connections map[net.Conn]bool) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accepted connection")
		connections[conn] = true
		// create conn
		tcp_client := connection.Stream{
			Type: "tcp",
			Conn: conn,
		}
		go message.Read(read_chan, tcp_client)
	}

}

func TCPClient() {
	address := flags.GetArg(3)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
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

func conn_cleanup(closed_conn chan net.Conn, conns map[net.Conn]bool) {
	for {
		conn := <-closed_conn
		conn.Close()
		delete(conns, conn)
	}

}

func write_to_conns(message_chan chan message.Message, connections map[net.Conn]bool) {
	for {
		message := <-message_chan
		// check message
		for conn := range connections {
			conn.Write(message.Content)
		}
	}
}
