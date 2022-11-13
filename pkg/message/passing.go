package message

import (
	"github.com/rs/zerolog/log"
)

// Read reads a message from the given reader
func Read(message_chan chan Message, reader Reader) {

	for {
		buf := reader.ReadMessage()
		if buf.Content != nil {
			message_chan <- buf
		} else {
			log.Error().Msg("could not read message")
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
