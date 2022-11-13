package protocol

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/rs/zerolog/log"
)

// UnixServer is a server for the unix socket, it bridged server to stdout and stdin to server
func UnixServer() {
	address := flags.Config.Host + ":" + strconv.Itoa(flags.Config.Port)
	listener, err := net.Listen("unix", address)
	if err != nil {
		log.Err(err).Msg("failed to listen on unix socket")
	}
	conn, err := listener.Accept()
	log.Info().Msg("Accepted connection")
	if err != nil {
		log.Err(err).Msg("failed to accept connection")
	}
	buf := make([]byte, 1024)
	for {
		_, err = conn.Read(buf)
		if err != nil {
			log.Err(err).Msg("failed to read from connection")
		}
		fmt.Println(string(buf))
	}

}

// UnixClient is a client for the unix socket, it bridged stdin to server and server to stdout
func UnixClient() {
	address := flags.Config.Host + ":" + strconv.Itoa(flags.Config.Port)
	conn, err := net.Dial("unix", address)
	if err != nil {
		log.Err(err).Msg("failed to dial unix socket")
	}
	defer conn.Close()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		} else {
			log.Error().Msg("failed to read from stdin")
		}
	}

}
