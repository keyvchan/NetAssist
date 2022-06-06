package connection

import (
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/message"
)

type Packet struct {
	Type string
	net.PacketConn
}

func (s Packet) ReadMessage() message.Message {
	return ReadPacketConn(s.PacketConn)
}

func ReadPacketConn(reader net.PacketConn) message.Message {
	buf := make([]byte, 1024)
	n, addr, err := reader.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	return message.Message{
		Content: buf[:n],
		Addr:    addr,
	}

}

// implements WriteMessage
func (s Packet) WriteMessage(msg message.Message) {
	WritePacketConn(s.PacketConn, msg)
}

func WritePacketConn(conn net.PacketConn, message message.Message) {
	_, err := conn.WriteTo(message.Content, message.Addr)
	if err != nil {
		log.Fatal(err)
	}
}
