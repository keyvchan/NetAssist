package client

import (
	"log"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/protocol"
)

func Req(args []string) {
	types := args[0]
	log.Println("Req:", types)
	switch types {
	case "tcp":
		protocol.TCPClient(args)
	case "udp":
		protocol.UDPClient(args)
	case "unix":
		protocol.UnixClient(args)
	case "unixgram":
		protocol.UnixgramClient(args)
	case "unixpacket":
		internal.Unimplemented("unixpacket")
	case "ip":
		internal.Unimplemented("ip")
	default:
		log.Fatal("unknown protocol", types)
	}
}
