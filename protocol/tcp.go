package protocol

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/internal"
)

func TCPServer() {
	address := internal.GetArg(3)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	// store all connections in a slice
	var connections = map[net.Conn]bool{}

	// create a chennel to communicate between the read and write goroutines
	message_chan := make(chan []byte)

	// create a goroutine to read user input
	go read_from_stdin(message_chan)
	go write_to_conns(message_chan, connections)

	// create a goroutines cleanup closed conn
	closed_conn := make(chan net.Conn)
	go conn_cleanup(closed_conn, &connections)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accepted connection")
		connections[conn] = true
		go read_from_conn(conn, closed_conn)
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
	message_chan := make(chan []byte)
	go read_from_stdin(message_chan)
	go write_to_conn(message_chan, conn)

	read_from_conn(conn, nil)

}

func conn_cleanup(closed_conn chan net.Conn, conns *map[net.Conn]bool) {
	for {
		conn := <-closed_conn
		conn.Close()
		delete(*conns, conn)
	}

}

func read_from_conn(conn net.Conn, closed_conn chan net.Conn) {
	buf := make([]byte, 1024)
	input_binary := internal.GetArg(4)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			log.Println("Connection closed")
			// remove from channel
			if closed_conn != nil {
				// server conn
				closed_conn <- conn
			} else {
				// client conn
				conn.Close()
			}
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(conn.RemoteAddr())
		if input_binary == "--binary" {
			for i := 0; i < n; i++ {
				fmt.Printf("%02x ", buf[i])
			}
			fmt.Print("\n")
		} else {
			fmt.Println(string(buf))
		}
	}
}

func write_to_conns(message_chan chan []byte, connections map[net.Conn]bool) {
	for {
		message := <-message_chan
		for conn := range connections {
			conn.Write(message)
		}
	}
}

func write_to_conn(message_chan chan []byte, conn net.Conn) {
	for {
		message := <-message_chan
		conn.Write(message)
	}
}

func read_from_stdin(message_chan chan []byte) {
	for {
		buf := internal.StdinRead()
		if buf != nil {
			message_chan <- buf
		} else {
			log.Fatal("Could not read from stdin")
		}
	}
}
