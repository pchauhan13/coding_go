package codinggo

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// Program16 : program to define a function using recursion and then using it
func Program16() {
	fmt.Println(fact(5))
}
