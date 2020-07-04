package codinggo

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// Program13 : program to use of function with multiple values
func Program13() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
