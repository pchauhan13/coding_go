package codinggo

import "fmt"

// Program24 : program to define and use buffered channels
func Program24() {
	messages := make(chan string, 2)

	messages <- "message1"
	messages <- "message2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
