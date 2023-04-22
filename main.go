package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				fmt.Println("hi bro")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for the goroutine to finish
	wg.Wait()
}
