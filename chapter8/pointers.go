package main

import "fmt"
/*
	Pointers reference a location in memory where a value is stored rather than the value itself. 
	(They point to something else) By using a pointer (*int) the zero function is able to modify 
	the original variable.

	In this program the zero function will not modify the original x variable in the main function.

		func zero(x int) { 
			x=0
		}

		func main() {
			x := 5
			zero(x)
			fmt.Println(x) // x is still 5
		}


	if we want to modify the original value, we must use a pointers

	In Go a pointer is represented using the `*` (asterisk) character followed by 
	the type of the stored value.

	`*` is also used to “dereference” pointer variables. 
	Dereferencing a pointer gives us access to the value the pointer points to.

	use the `&` operator to find the address of a variable.
*/

func zero(xPtr *int) {
	//	When we write *xPtr = 0 we are saying 
	// “store the `int` 0 in the memory location `xPtr` refers to”.
	*xPtr = 0
}

func main() {
	x := 5
	//	we use the `&` operator to find the address of a variable. 
	//	`&x` returns a `*int` (pointer to an `int`) because `x` is an `int`.
	zero(&x)	//	<- `&` operator used to find the address of a variable
	fmt.Println(x) // x is 0
}
