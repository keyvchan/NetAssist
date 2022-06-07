package message

import (
	"errors"
	"log"
)

// Read reads a message from the given reader
func Read(message_chan chan Message, reader Reader) {

	for {
		buf := reader.ReadMessage()
		if buf.Content != nil {
			message_chan <- buf
		} else {
			log.Println(errors.New("could not read message"), reader)
		}
	}

}

func Write(message_chan chan Message, writer Writer) {
	for {
		// receive message from message_chan
		buf := <-message_chan
		writer.WriteMessage(buf)
	}
}
