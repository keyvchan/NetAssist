package internal

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net"
	"os"
)

func ReadStdin(reader interface{}) Message {
	stdin, ok := reader.(*os.File)
	if !ok {
		log.Fatal("Wrong type")
	}
	input_binary := GetArg(4)
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

		return Message{
			Content: buf,
			Conn:    nil,
		}
	} else {
		log.Fatal(errors.New("failed to read from stdin"))
	}
	return Message{}
}

var ClosedConn = new(chan net.Conn)

func ReadConn(conn interface{}) Message {
	// type checking
	connn, ok := conn.(net.Conn)
	if !ok {
		log.Fatal("Wrong type")
	}
	buf := make([]byte, 1024)
	// input_binary := GetArg(4)
	n, err := connn.Read(buf)
	if errors.Is(err, io.EOF) {
		log.Println("Connection closed")
		// remove from channel
		*ClosedConn <- connn
		return Message{}
	}
	if err != nil {
		log.Println(err)
		return Message{}
	}
	message := Message{
		Conn:    connn,
		Content: buf[:n],
	}
	return message
}

func ReadPacketConn(reader interface{}) Message {
	conn, ok := reader.(net.PacketConn)
	if !ok {
		log.Fatal("Wrong type")
	}

	buf := []byte{}
	_, _, err := conn.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	return Message{
		Content: buf,
		Conn:    nil,
	}

}
