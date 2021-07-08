package main

import "fmt"

var y = 22

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is a TYPE string
// and I ASSIGN the VALUE "shaken not stirred"

var z = "shaken not stirred"
var a = `James said "Shaken"`

// this is a STATIC programmin language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	//Print the type of VARIABLE
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
