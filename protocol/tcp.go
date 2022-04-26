package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
)

// store all connections in a slice
var connections = map[net.Conn]bool{}

func TCPServer() {
	address := internal.GetArg(3)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	// create a chennel to communicate between the read and write goroutines
	message_chan := make(chan []byte)

	// create a goroutine to read user input
	go read_from_stdin(message_chan)
	go write_to_conn(message_chan)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accepted connection")
		connections[conn] = true
		go func() {
			buf := make([]byte, 1024)
			for {
				_, err = conn.Read(buf)
				if errors.Is(err, io.EOF) {
					log.Println("Connection closed")
					// remove from channel
					delete(connections, conn)
					break
				}
				if err != nil {
					log.Println(err)
				}
				fmt.Println(conn.RemoteAddr())
				fmt.Println(string(buf))
			}

		}()
	}
}

func write_to_conn(message_chan chan []byte) {
	for {
		message := <-message_chan
		for conn := range connections {
			conn.Write(message)
		}
	}

}

func read_from_stdin(message_chan chan []byte) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message_chan <- scanner.Bytes()
		} else {
			log.Fatal(errors.New("failed to read from stdin"))
		}
	}
}

func TCPClient() {
	address := internal.GetArg(3)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		} else {
			log.Fatal(errors.New("failed to read from stdin"))
		}

	}

}
