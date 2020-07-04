package codinggo

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Program18 : program to define and use structs
func Program18() {
	fmt.Println(person{"Bob", 42})
	fmt.Println(person{name: "Alice", age: 44})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 20})
	fmt.Println(newPerson("Prashant"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 44
	fmt.Println(s.age)
}
