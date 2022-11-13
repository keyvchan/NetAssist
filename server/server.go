package server

import (
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol"
	"github.com/rs/zerolog/log"
)

func Serve() {
	types := flags.Config.Protocol
	log.Info().Msg("Serve: " + types)
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
		log.Error().Msg("unknow protocol: " + types)
	}
}
