package main

import "fmt"

func main() {
	c := make(chan string)
	go SayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}
	v, ok := <-c
	fmt.Println("Channel close?", !ok, " Value: ", v)
}

func SayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- "Hello"
	}
	close(c)
}
