package main

import "fmt"

func main() {
	fmt.Println("It's a line break output.")
	fmt.Printf("It's a %s output.\n", "format")
	fmt.Print("It's a general output.\n")
	fmt.Println(fmt.Sprintf("%s resulting a format string as output.", "Sprintf"))
	// fmt.Fprint(new(bytes.Buffer), "Hello to the writer interface!")
}
