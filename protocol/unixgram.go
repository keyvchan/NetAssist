package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/internal"
)

func UnixgramServer() {
	address := internal.GetArg(3)
	conn, err := net.ListenPacket("unixgram", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		conn.ReadFrom(buf)
		fmt.Println(string(buf))
	}
}

func UnixgramClient() {
	address := internal.GetArg(3)
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
