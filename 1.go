package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var c int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			c++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Просмотров ", c)
}
