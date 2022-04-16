package server

import (
	"log"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/protocol"
)

func Serve() {
	types := internal.GetArg(2)
	log.Println("Serve:", types)
	switch types {
	case "tcp":
		protocol.TCPServer()
	case "udp":
		protocol.UDPServer()
	case "unix":
		protocol.UnixServer()
	case "unixgram":
		protocol.UnixgramServer()
	case "unixpacket":
		internal.Unimplemented("unixpacket")
	case "ip":
		internal.Unimplemented("ip")
	default:
		log.Fatal("unknow protocol: ", types)
	}
}
