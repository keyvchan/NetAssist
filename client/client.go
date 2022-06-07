package client

import (
	"log"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol/tcp"
	"github.com/keyvchan/NetAssist/protocol/udp"
	"github.com/keyvchan/NetAssist/protocol/unixgram"
	"github.com/keyvchan/NetAssist/protocol/unixsocket"
)

// Req is the entry point for the client
func Req() {
	types := flags.GetArg(2)
	log.Println("Req:", types)
	switch types {
	case "tcp":
		tcp.Client()
	case "udp":
		udp.Client()
	case "unix":
		unixsocket.Client()
	case "unixgram":
		unixgram.Client()
	case "unixpacket":
		utils.Unimplemented("unixpacket")
	case "ip":
		utils.Unimplemented("ip")
	default:
		log.Fatal("unknown protocol", types)
	}
}
