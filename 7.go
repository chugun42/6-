package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	stock := 10

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			if stock > 0 {
				stock--
				fmt.Println("Товар продан, остаток:", stock)
			}
			mu.Unlock()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			stock++
			fmt.Println("Товар поступил, остаток:", stock)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый остаток:", stock)
}
