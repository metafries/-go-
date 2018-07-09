package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)        // Increment the WaitGroup counter
		go func(i int) { // Launch a goroutine
			defer wg.Done() // Decrement the counter when the goroutine completes
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for ", i)
		}(i)
	}
	wg.Wait()
}
