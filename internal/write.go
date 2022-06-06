package internal

import (
	"fmt"
	"log"
	"net"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
)

func WriteStdout(writter interface{}, message message.Message) {

	// write to stdout
	input_binary := flags.GetArg(4)
	fmt.Println(message.Addr)

	if input_binary == "--binary" {
		for i := 0; i < len(message.Content); i++ {
			fmt.Printf("%02x ", message.Content[i])
		}
		fmt.Print("\n")
	} else {
		fmt.Println(string(message.Content))
	}
}

func WriteConn(writter interface{}, message message.Message) {
	conn, ok := writter.(net.Conn)
	if !ok {
		log.Fatal("Wrong type")
	}
	_, err := conn.Write(message.Content)
	if err != nil {
		log.Fatal(err)
	}
}
