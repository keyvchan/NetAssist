package message

import (
	"net"
)

type Message struct {
	Content []byte
	net.Addr
}

type Reader interface {
	ReadMessage() Message
}

type Writer interface {
	WriteMessage(message Message)
}
