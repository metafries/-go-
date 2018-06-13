package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}("World")
	fmt.Print("Hello")
	time.Sleep(3 * time.Second)
}
