package main

import "fmt"

type Set map[string]struct{}

func getSetKeys(s Set) []string {
	var rv []string
	for k := range s {
		rv = append(rv, k)
	}
	return rv
}

func main() {
	s := make(Set)
	s["item_1"] = struct{}{}
	s["item_2"] = struct{}{}
	fmt.Println(getSetKeys(s))
}
