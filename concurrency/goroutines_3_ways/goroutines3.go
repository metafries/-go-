package main

import (
	"fmt"
	"time"
)

func main() {
	word := "Hello"
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()
	fmt.Print(word)
	word = "World"
	time.Sleep(3 * time.Second)
}
