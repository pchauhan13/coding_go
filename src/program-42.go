package codinggo

import "os"

// Program42 : program to use panic in programs
func Program42() {
	panic("a problem happened")

	_, err := os.Create("./file")
	if err != nil {
		panic(err)
	}
}
