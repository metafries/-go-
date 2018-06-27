package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RangeRand(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) GetCount() int {
	return cr.count
}

func main() {
	cr := NewCustomRand(time.Now().UnixNano())
	fmt.Println(cr.RangeRand(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())
}
