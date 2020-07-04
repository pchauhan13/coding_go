package codinggo

import (
	"fmt"
	"time"
)

// Program26 : program to use select to wait upon multiple channels
func Program26() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "done 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "done 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}
