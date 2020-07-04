package codinggo

import "fmt"

func sum(values ...int) {
	fmt.Println(values, " ")
	total := 0
	for _, num := range values {
		total += num
	}
	fmt.Println(total)
}

// Program15 : program to make use of variadic functions
func Program15() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
