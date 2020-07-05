package codinggo

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Program37 : program to create atomic counters for goroutines synchronization
func Program37() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
