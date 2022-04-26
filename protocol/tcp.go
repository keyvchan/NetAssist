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

func TCPServer() {
	address := internal.GetArg(3)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		go func() {
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
				fmt.Println(conn.RemoteAddr())
				fmt.Println(string(buf))
			}

		}()
	}
}
func TCPClient() {
	address := internal.GetArg(3)
	conn, err := net.Dial("tcp", address)
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
