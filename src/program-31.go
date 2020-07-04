package codinggo

import "fmt"

// Program31 : program to use range over channels
func Program31() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
