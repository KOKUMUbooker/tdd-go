package main

import "fmt"

// Naming structure for "Example" + name of function
func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
