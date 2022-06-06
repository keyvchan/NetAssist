package message

import (
	"log"
)

// Read reads a message from the given reader
func Read(message_chan chan Message, reader Reader) {

	for {
		buf := reader.ReadMessage()
		if buf.Content != nil {
			message_chan <- buf
		} else {
			log.Fatal("Could not read from conn")
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
