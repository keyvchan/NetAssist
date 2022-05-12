package internal

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"log"
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
