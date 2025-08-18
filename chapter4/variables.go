package main

import "fmt"

/*
	A variable is a storage location, 
	with a specific type and an associated name.

	The type is not necessary because the Go compiler 
	is able to infer the type based on the literal value 
	you assign the variable.

	x := "Hello World"

	Since you are assigning a string literal, 
	x is given the type string

	var x = "Hello World"
	var x string = "Hello World"

	Naming variables:
	Names must start with a letter and may contain letters, 
	numbers or the _ (underscore) symbol.
*/

func main() {
	var x string	//	variable
	x = "first "
	fmt.Println(x)
	
	x = x + "second"
	//	alternative: x += "second"
	fmt.Println(x)

	// var x string = "hello"
	// var y string = "world"
	// fmt.Println(x == y)
	// 	output: false

	// var x string = "hello"
	// var y string = "hello"
	// fmt.Println(x == y)
	//	output: true
}

/*
alternative:

package main

import "fmt"

func main() {
    var x string
    x = "Hello World"
    fmt.Println(x)
}

*/

/*

short form:
x := "Hello World"
same as:
var x = "Hello World"

another short form:
x := 5
fmt.Println(x)


Generally you should use this shorter form whenever possible.
*/