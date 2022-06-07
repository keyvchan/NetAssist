package utils

import (
	"errors"
	"log"
)

// Unimplemented is a stub function that returns an error
func Unimplemented(message string) {
	log.Fatal(errors.New("unimplemented: " + message))
}
