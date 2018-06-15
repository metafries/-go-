package main

import "fmt"

func main() {
	c := make(chan bool)
	go waitAndSay(c, "World")
	fmt.Print("Hello")
	c <- true // Send a signal to c in order to allow waitAndSay to continue
	<-c       // Wait and receive another signal on c before exit
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}
