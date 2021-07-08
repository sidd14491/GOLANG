/* 1. Use a var to DECLARE a three variables. The variable should have package
	level scope. Do not assign VALUES to the variables . USE the following
	IDENTIFIERS for the variables and make sure the variables are of the following
	TYPE
		a. identifier "x" type int
		b. identifier "y" type string
		c. identifier "z" type bool
2. in func main
	a. print out the value of each identifier
	b.The compiler assigned values to the variables. What are these values called
*/
package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\t%v\n", x, x)
	fmt.Println("String value", y, "ending")
	fmt.Printf("%T\t%v\n", z, z)
}
