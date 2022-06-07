package client

import (
	"log"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol"
)

// Req is the entry point for the client
func Req() {
	types := flags.GetArg(2)
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
		utils.Unimplemented("unixpacket")
	case "ip":
		utils.Unimplemented("ip")
	default:
		log.Fatal("unknown protocol", types)
	}
}
