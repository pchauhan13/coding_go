package codinggo

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// Program27 : program to
func Program27() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
