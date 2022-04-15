package internal

import "log"

func Unimplemented(things ...interface{}) {
	log.Fatal("unimplemented", things)
}
