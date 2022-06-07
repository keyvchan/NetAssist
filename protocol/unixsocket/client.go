package unixsocket

import (
	"bufio"
	"errors"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/pkg/flags"
)

// UnixClient is a client for the unix socket, it bridged stdin to server and server to stdout
func Client() {
	address := flags.GetArg(3)
	conn, err := net.Dial("unix", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			conn.Write([]byte(scanner.Text()))
		} else {
			log.Fatal(errors.New("failed to read from stdin"))
		}
	}

}
