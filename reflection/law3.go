package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Sample_1
	var x1 float32 = 5.8
	fmt.Println("x1 -> InitVal:", x1)
	v1p := reflect.ValueOf(&x1)
	fmt.Println("Is x1(v1p) settable?", v1p.CanSet())
	v1pElem := v1p.Elem()
	fmt.Println("Is x1(v1pElem) settable?", v1pElem.CanSet())
	v1pElem.SetFloat(2.2)
	fmt.Println("x1 -> SetVal:", x1)
	// Sample_2
	x2 := 256
	fmt.Println("x2 -> InitVal:", x2)
	DoubleInt(&x2)
	fmt.Println("x2 -> DoubleInt(x2):", x2)
}

func DoubleInt(n *int) {
	*n = *n * 2
}
