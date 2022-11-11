package perf

import (
	"log"
	"os"
)

func ServerPerf() {
	log.Println("perf")

	// check os.arg
	if os.Args[2] == "concurrency" {
		perf_concurrency()
	}

}
