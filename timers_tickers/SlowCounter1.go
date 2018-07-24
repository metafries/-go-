package main

import (
	"fmt"
	"time"
)

func main() {
	go SlowCounter(2)
	time.Sleep(15 * time.Second)
}

func SlowCounter(n int) {
	i := 0
	d := time.Duration(n) * time.Second // Create a duration of n seconds
	for {                               // Create a timer for this duration
		t := time.NewTimer(d)
		<-t.C
		i++
		fmt.Println(i)
	}
}
