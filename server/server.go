package server

import (
	"log"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/protocol"
)

func Serve(args []string) {
	types := args[0]
	log.Println("Serve:", types)
	switch types {
	case "tcp":
		protocol.TCPServer(args)
	case "udp":
		protocol.UDPServer(args)
	case "unix":
		protocol.UnixServer(args)
	case "unixpacket":
		internal.Unimplemented("unixpacket")
	case "ip":
		internal.Unimplemented("ip")
	default:
		log.Fatal("unknow protocol: ", types)
	}
}
