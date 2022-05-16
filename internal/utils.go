package internal

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Unimplemented(things ...interface{}) {
	log.Fatal("unimplemented", things)
}

func StdinRead() []byte {
	input_binary := GetArg(4)
	scanner := bufio.NewScanner(os.Stdin)
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
		return buf
	} else {
		log.Fatal(errors.New("failed to read from stdin"))
	}
	return nil
}

func ConnRead(conn net.Conn, closed_conn chan net.Conn) {
	buf := make([]byte, 1024)
	input_binary := GetArg(4)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			log.Println("Connection closed")
			// remove from channel
			if closed_conn != nil {
				// server conn
				closed_conn <- conn
			} else {
				// client conn
				conn.Close()
			}
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(conn.RemoteAddr())
		if input_binary == "--binary" {
			for i := 0; i < n; i++ {
				fmt.Printf("%02x ", buf[i])
			}
			fmt.Print("\n")
		} else {
			fmt.Println(string(buf))
		}
	}
}
