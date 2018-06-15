package main

import (
	"fmt"
	"math/rand"
	"time"
)

var scMapping = map[string]int{
	"RUS-KSA": 5,
	"EGY-URU": 1,
}

func findSC(game, server string, c chan int) {
	time.Sleep(time.Duration(rand.Intn(120)) * time.Second) // Simulate searching
	c <- scMapping[game]                                    // Return security clearance from map
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c1 := make(chan int)
	c2 := make(chan int)
	game := "EGY-URU"
	go findSC(game, "server 1", c1)
	go findSC(game, "server 2", c2)
	select {
	case sc := <-c1:
		fmt.Println(game, " has a security clearance of ", sc, " found in server 1")
	case sc := <-c2:
		fmt.Println(game, " has a security clearance of ", sc, " found in server 2")
	case <-time.After(2 * time.Minute):
		fmt.Println("Timeout!")
	}
}
