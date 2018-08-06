package main

import (
	"-go-/_mfConfig"
	"fmt"
)

type ConfS struct {
	TS      string  `name:"testString" json:"testString"`
	TB      bool    `name:"testBool" json:"testBool"`
	TF      float64 `name:"testFloat" json:"testFloat"`
	TestInt int
}

func main() {
	cs := new(ConfS)
	// _mfConfig.GetConfiguration(_mfConfig.CUSTOM, cs, "test.conf")
	_mfConfig.GetConfiguration(_mfConfig.JSON, cs, "test.json")
	fmt.Println(*cs)
	if cs.TB {
		fmt.Println("bool is true")
	}
	fmt.Println(float64(4.8 * cs.TF))
	fmt.Println(5 * cs.TestInt)
	fmt.Println(cs.TS)
}
