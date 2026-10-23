package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			fmt.Println("Горутина", i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
}
