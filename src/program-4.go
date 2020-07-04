package codinggo

import (
	"fmt"
	"math"
)

const s string = "constant"

// Program4 : program to define and use constants
func Program4() {
	fmt.Println(s)

	const n = 5000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
