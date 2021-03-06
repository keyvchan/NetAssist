package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/keyvchan/NetAssist/pkg/flags"
)

// UnixServer is a server for the unix socket, it bridged server to stdout and stdin to server
func UnixServer() {
	address := flags.GetArg(3)
	listener, err := net.Listen("unix", address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	log.Println("Accepted connection")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	}

}

// UnixClient is a client for the unix socket, it bridged stdin to server and server to stdout
func UnixClient() {
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
