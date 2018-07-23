package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

// CLI: go run -race mapCounter.go
func main() {
	mc := MapCounter{
		m: make(map[int]int),
	}
	go runWriters(&mc, 10)
	go runReaders(&mc, 10)
	go runReaders(&mc, 10)
	time.Sleep(15 * time.Second)
}

func runWriters(mc *MapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReaders(mc *MapCounter, n int) {
	for {
		mc.RLock()
		v := mc.m[rand.Intn(n)] // [0] till [n-1]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}
