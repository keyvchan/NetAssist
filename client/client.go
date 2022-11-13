package client

import (
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/utils"
	"github.com/keyvchan/NetAssist/protocol"
	"github.com/rs/zerolog/log"
)

// Req is the entry point for the client
func Req() {
	types := flags.Config.Protocol
	log.Info().Msg("Req: " + types)
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
		log.Error().Msg("unknown protocol " + types)
	}
}
