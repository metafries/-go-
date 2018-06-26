package main

import "fmt"

type store struct {
	value interface{}
	name  string
}

func (s *store) SetValue(v interface{}) {
	s.value = v
}

func (s *store) GetValue() interface{} {
	return s.value
}

func NewStore(nm string) *store {
	return &store{
		name: nm,
	}
}

func main() {
	istore := NewStore("Interger Store")
	istore.SetValue(4)
	if v, ok := istore.GetValue().(int); ok {
		v *= 4
		fmt.Println(v)
	}
	sstore := NewStore("String Store")
	sstore.SetValue("Hello")
	if v, ok := sstore.GetValue().(string); ok {
		v += " World"
		fmt.Println(v)
	}
}
