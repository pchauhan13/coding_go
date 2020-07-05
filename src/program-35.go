package codinggo

import (
	"fmt"
	"sync"
	"time"
)

func workerForWait(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting...\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done.\n", id)
}

// Program35 : program to use waitgroups for multiple goroutines
func Program35() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerForWait(i, &wg)
	}

	wg.Wait()
}
