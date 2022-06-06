package connection

import (
	"errors"
	"io"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/message"
)

type Stream struct {
	Type string
	net.Conn
}

// implements ReadMessage for both type
func (s Stream) ReadMessage() message.Message {
	return ReadConn(s.Conn)
}

var ClosedConn = new(chan net.Conn)

func ReadConn(conn interface{}) message.Message {
	// type checking
	connn, ok := conn.(net.Conn)
	if !ok {
		log.Fatal("Wrong type")
	}
	buf := make([]byte, 1024)
	// input_binary := GetArg(4)
	n, err := connn.Read(buf)
	if errors.Is(err, io.EOF) {
		log.Println("Connection closed")
		// remove from channel
		*ClosedConn <- connn
		return message.Message{}
	}
	if err != nil {
		log.Println(err)
		return message.Message{}
	}
	message := message.Message{
		Addr:    connn.RemoteAddr(),
		Content: buf[:n],
	}
	return message
}

// implements WriteMessage
func (s Stream) WriteMessage(msg message.Message) {
	WriteConn(s.Conn, msg)
}

func WriteConn(conn net.Conn, message message.Message) {
	_, err := conn.Write(message.Content)
	if err != nil {
		log.Fatal(err)
	}
}
