package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	votes := make(map[string]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			votes["Кандидат А"]++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Голосов за А:", votes["Кандидат А"])
}
