package main

import (
	"fmt"
	"time"
)

func main() {
	go waitAndSay("World")
	fmt.Print("Hello")
	time.Sleep(3 * time.Second)
}

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}
