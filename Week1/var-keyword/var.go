package main

import "fmt"

// DECLARE the variable y
// ASSIGN the value 43
// declare and assign = initialization
var y = 43

// DECLARE there is variable with IDENTIFIER "z"
// and that the variable with the IDENTIFIER "z" if of TYPE int
// ASSIGN the ZERO VALUE of TYPE int to "z"
// "false" for booleans, 0 for int, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels and maps
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain type)
	x := 12
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println(y)
}
