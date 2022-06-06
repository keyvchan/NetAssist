package connection

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/pkg/message"
)

// File represents a abstraced file connection.
type File struct {
	Path string
	File *os.File
}

// pre-initialize the file connection
var Stdin = File{
	Path: "/dev/stdin",
	File: os.Stdin,
}

var Stdout = File{
	Path: "/dev/stdout",
	File: os.Stdout,
}

// implements ReadMessage
func (f File) ReadMessage() message.Message {
	switch f.Path {
	case "/dev/stdin":
		return ReadStdin(f.File)

	default:
		return message.Message{}
	}
}

func ReadStdin(stdin *os.File) message.Message {
	input_binary := flags.GetArg(4)
	scanner := bufio.NewScanner(stdin)
	if scanner.Scan() {

		buf := []byte{}
		if input_binary == "--binary" {
			byte_slices := bytes.Split(scanner.Bytes(), []byte(" "))
			for _, byte_slice := range byte_slices {
				new_byte := make([]byte, 1)
				_, err := hex.Decode(new_byte, byte_slice)
				if err != nil {
					log.Println(err)
				}
				buf = append(buf, new_byte...)
			}

		} else {
			buf = scanner.Bytes()
		}
		// split by space

		return message.Message{
			Content: buf,
			Addr:    nil,
		}
	} else {
		log.Fatal(errors.New("failed to read from stdin"))
	}
	return message.Message{}
}

func (f File) WriteMessage(msg message.Message) {
	switch f.Path {
	case "/dev/stdout":
		WriteStdout(f.File, msg)
	default:
		return
	}
}

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
