package codinggo

import (
	"fmt"
	"sort"
)

// Program40 : program to implement sorting with custom sorting fucntions
func Program40() {
	str := []string{"c", "a", "b"}
	sort.Strings(str)
	fmt.Println("Strings:", str)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
