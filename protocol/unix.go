package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

func UnixServer(args []string) {
	address := args[1]
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

func UnixClient(args []string) {
	address := args[1]
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
