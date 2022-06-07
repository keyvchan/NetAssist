package tcp

import (
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/connection"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
)

// TCPClient is a TCP client, read from stdin and write to the server and read from the server when it to stdout
func Client() {
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
