package internal

import (
	"log"
	"net"
	"os"
)

var args []string

func SetArgs() {
	args = os.Args
}

func GetArg(i int) string {
	if i < len(args) {
		return args[i]
	} else {
		if i == 4 {
			return "text"
		}
		log.Fatal("Index out of range")
	}
	return ""
}

type Message struct {
	Content []byte
	Conn    net.Conn
}
