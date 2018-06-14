package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"
	fmt.Print(<-ch)
	fmt.Println(<-ch)
}
