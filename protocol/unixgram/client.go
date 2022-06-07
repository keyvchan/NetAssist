package unixgram

import (
	"bufio"
	"errors"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/pkg/flags"
)

// Client is a client for the unixgram protocol, its create a bridge between server and client
func Client() {
	address := flags.GetArg(3)
	conn, err := net.Dial("unixgram", address)
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
