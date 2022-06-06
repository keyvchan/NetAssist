package utils

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/pkg/message"
)

// ReadMessage reads a message from the given reader
func ReadMessage(message_chan chan message.Message, reader message.Reader) {
	// var source_func func(interface{}) internal.Message
	// switch reader.(type) {
	// case net.Conn:
	// 	source_func = internal.ReadConn
	// case *os.File:
	// 	source_func = internal.ReadStdin
	// case net.PacketConn:
	// 	source_func = internal.ReadPacketConn
	// }

	for {
		buf := reader.ReadMessage()
		if buf.Content != nil {
			message_chan <- buf
		} else {
			log.Fatal("Could not read from conn")
		}
	}

}

func WriteMessage(message_chan chan message.Message, writer interface{}) {
	var dest_func func(interface{}, message.Message)
	switch writer.(type) {
	case net.Conn:
		dest_func = internal.WriteConn
	case *os.File:
		dest_func = internal.WriteStdout
	}
	for {
		// receive message from message_chan
		buf := <-message_chan
		dest_func(writer, buf)
	}
}
