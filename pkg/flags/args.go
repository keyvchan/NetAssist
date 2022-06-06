package flags

import (
	"log"
	"os"
)

// args stores the arguments
var args []string

// SetArgs retrieves the arguments from the command line
func SetArgs() {
	args = os.Args
}

// GetArg returns the nth argument
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
