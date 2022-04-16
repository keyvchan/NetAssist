package client

import (
	"log"

	"github.com/keyvchan/NetAssist/internal"
	"github.com/keyvchan/NetAssist/protocol"
	"github.com/keyvchan/NetAssist/protocol/ip"
)

func Req() {
	types := internal.GetArg(2)
	log.Println("Req:", types)
	switch types {
	case "tcp":
		protocol.TCPClient()
	case "udp":
		protocol.UDPClient()
	case "unix":
		protocol.UnixClient()
	case "unixgram":
		protocol.UnixgramClient()
	case "unixpacket":
		internal.Unimplemented("unixpacket")
	case "icmp":
		ip.ICMPRequest()
	default:
		log.Fatal("unknown protocol", types)
	}
}
