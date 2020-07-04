package codinggo

import "fmt"

// Program10 : program to define and use maps
func Program10() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 12

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len of map:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
