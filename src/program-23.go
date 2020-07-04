package codinggo

import "fmt"

// Program23 : program to define channels and use them
func Program23() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
