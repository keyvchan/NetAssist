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

func ReadPacketConn(reader interface{}) message.Message {
	conn, ok := reader.(net.PacketConn)
	if !ok {
		log.Fatal("Wrong type")
	}

	buf := []byte{}
	_, _, err := conn.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	return message.Message{
		Content: buf,
		Addr:    nil,
	}

}
