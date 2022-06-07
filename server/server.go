package server

import (
	"log"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol"
)

func Serve() {
	types := flags.GetArg(2)
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
		utils.Unimplemented("unixpacket")
	case "ip":
		utils.Unimplemented("ip")
	default:
		log.Fatal("unknow protocol: ", types)
	}
}
