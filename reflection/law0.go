package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	c, _ := net.Dial("tcp", ":2300")
	var r io.Reader
	r = c // r now stores (value:c, type:Conn)
	// That's why we can also do this:
	if _, ok := r.(io.Writer); ok {
		/*
			Even though r in theory is only of type io.Reader,
			the underlying value stored also implements the io.Writer interface.
		*/
		fmt.Println("There is also a writer inside value c!")
	}
}
