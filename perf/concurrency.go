package perf

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func perf_concurrency() {
	fmt.Println("perf_concurrency")
	// TODO
	// create multiple goroutines
	pool := make(chan bool, 100)

	for {
		pool <- true
		go func(pool chan bool) {
			conn, err := net.Dial("tcp", os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			log.Println("connected")
			// alive for random time

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			<-pool
		}(pool)
	}

}
