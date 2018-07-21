package main

import (
	"fmt"
	"reflect"
)

type customFloat float64

func main() {
	var x1 float32 = 5.7
	inspectIfTypeFloat(x1)
	var x2 customFloat = 5.7
	fmt.Println("Is x2 customFloat type same as float64?", reflect.ValueOf(x2).Kind() == reflect.Float64)
	fmt.Println("x2 -> type / value: ", reflect.ValueOf(x2).Type(), "/", x2)
}

func inspectIfTypeFloat(i interface{}) {
	p := reflect.ValueOf(i)
	fmt.Println("p -> type / value: ", p.Type(), "/", p)
	fmt.Println("p.Float() -> type / value: ", reflect.TypeOf(p.Float()), "/", p.Float())
}
