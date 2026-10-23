package main

import (
	"fmt"
	"sync"
)

var (
	mu   sync.RWMutex
	data = make(map[string]string)
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		data["t"] = "текст"
		mu.Unlock()
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			fmt.Println("Чтение:", data["t"])
			mu.RUnlock()
		}()
	}

	wg.Wait()
}
