package codinggo

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// Program25 : program to use channel synchronization
func Program25() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
