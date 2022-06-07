package server

import (
	"log"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol/tcp"
	"github.com/keyvchan/NetAssist/protocol/udp"
	"github.com/keyvchan/NetAssist/protocol/unixgram"
	"github.com/keyvchan/NetAssist/protocol/unixsocket"
)

func Serve() {
	types := flags.GetArg(2)
	log.Println("Serve:", types)
	switch types {
	case "tcp":
		tcp.Server()
	case "udp":
		udp.Server()
	case "unix":
		unixsocket.Server()
	case "unixgram":
		unixgram.Server()
	case "unixpacket":
		utils.Unimplemented("unixpacket")
	case "ip":
		utils.Unimplemented("ip")
	default:
		log.Fatal("unknow protocol: ", types)
	}
}
