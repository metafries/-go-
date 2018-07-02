package main

import (
	"-go-/factory/appliances"
	"fmt"
)

func main() {
	// Request the user to enter the appliance type
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	// Use fmt.scan to retrieve the user's input
	var input int
	fmt.Scan(&input)

	appliance, err := appliances.CreateAppliance(input)

	// If no errors, start the appliance then print it's purpose
	if err == nil {
		appliance.Start()
		fmt.Println(appliance.GetPurpose())
		// If error encountered, print the error
	} else {
		fmt.Println(err)
	}
}
