package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"TIME"`
		Field2 string  `alias:"f2" desc:"CLUBvCLUB"`
		Field3 float64 `alias:"f3" desc:"DATE"`
	}
	mys := myStruct{1005, "ROMAvTOT", 7.26}
	InspectStructType(&mys)
}

func InspectStructType(i interface{}) {
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr {
		return
	}
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return
	}
	mysRValue.Field(0).SetInt(1105)
	mysRValue.Field(1).SetString("BARvTOT")
	mysRValue.Field(2).SetFloat(7.29)
	mysRType := mysRValue.Type()
	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("Name: %s | Type: %s | Value: %v | Tag(alias): %s | Tag(desc): %s \n",
			fieldRType.Name, fieldRType.Type, fieldRValue.Interface(), fieldRType.Tag.Get("alias"), fieldRType.Tag.Get("desc"))

	}
}
