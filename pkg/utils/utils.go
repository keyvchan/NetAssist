package utils

import (
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
)

func ReadMessage(message_chan chan internal.Message, reader interface{}) {
	var source_func func(interface{}) internal.Message
	switch reader.(type) {
	case net.Conn:
		source_func = internal.ReadConn
	case *os.File:
		source_func = internal.ReadStdin
	case net.PacketConn:
		source_func = internal.ReadPacketConn
	}

	for {
		buf := source_func(reader)
		if buf.Content != nil {
			message_chan <- buf
		} else {
			log.Fatal("Could not read from conn")
		}
	}

}

func WriteMessage(message_chan chan internal.Message, writer interface{}) {
	var dest_func func(interface{}, internal.Message)
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
