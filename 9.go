package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	success, fails := 0, 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			if i%2 == 0 {
				success++
			} else {
				fails++
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Успешных: %d, Ошибок: %d\n", success, fails)
}
