package utils

import (
	"errors"
	"log"
)

func Unimplemented(message string) {
	log.Fatal(errors.New("unimplemented: " + message))
}
