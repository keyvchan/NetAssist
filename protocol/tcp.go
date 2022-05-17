package protocol

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/pkg/utils"
)

func TCPServer() {
	address := internal.GetArg(3)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	// store all connections in a slice
	// NOTE: Possibly race condition
	var connections = map[net.Conn]bool{}

	// create a chennel to communicate between the read and write goroutines
	conn_read := make(chan internal.Message)
	stdin_read := make(chan internal.Message)

	// create a goroutine to read user input
	go utils.ReadMessage(stdin_read, os.Stdin)
	go utils.WriteMessage(conn_read, os.Stdout)

	// create a goroutines cleanup closed conn
	go conn_cleanup(*internal.ClosedConn, connections)
	go write_to_conns(stdin_read, connections)
	go accept_conn(conn_read, listener, connections)
	select {}

}

func accept_conn(read_chan chan internal.Message, listener net.Listener, connections map[net.Conn]bool) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accepted connection")
		connections[conn] = true
		go utils.ReadMessage(read_chan, conn)
	}

}

func TCPClient() {
	address := internal.GetArg(3)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create a chennel to communicate between the read and write goroutines
	stdin_read := make(chan internal.Message)
	go utils.ReadMessage(stdin_read, os.Stdin)
	go utils.WriteMessage(stdin_read, conn)

	quit := make(chan bool)
	conn_read := make(chan internal.Message)
	go utils.ReadMessage(conn_read, conn)
	go utils.WriteMessage(conn_read, os.Stdout)
	go func() {
		conn := <-*internal.ClosedConn
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

func write_to_conns(message_chan chan internal.Message, connections map[net.Conn]bool) {
	for {
		message := <-message_chan
		// check message
		for conn := range connections {
			conn.Write(message.Content)
		}
	}
}
