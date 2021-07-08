package main

import (
	"fmt"
)

var y string
var z int

func main() {
	// DECLARE a VARIABLE for certain TYPE
	// and the ASSIGN a VALUE of that type to the variable

	fmt.Println("Print the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "James bon James"
	fmt.Println("Print the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
