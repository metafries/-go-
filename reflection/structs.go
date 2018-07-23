package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int
		Field2 string
		Field3 float64
	}
	mys := myStruct{1005, "ROMAvTOT", 7.26}
	InspectStructType(mys)
}

func InspectStructType(i interface{}) {
	mysRValue := reflect.ValueOf(i)
	mysRType := reflect.TypeOf(i)
	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("Name: %s | Type: %s | Value: %v \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
	}
}
