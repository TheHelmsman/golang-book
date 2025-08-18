package main

import "fmt"

/*
	Another way to get a pointer is to use the built-in new function:
		
		func one(xPtr *int) {
				*xPtr = 1
		}
		
		func main() {
				xPtr := new(int)
				one(xPtr)
				fmt.Println(*xPtr) // x is 1
		}

	new takes a type as an argument, allocates enough memory to fit a 
	value of that type and returns a pointer to it.

	Pointers are rarely used with Go's built-in types, but they are 
	extremely useful when paired with structs.
*/

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	//	Another way to get a pointer is to use the built-in new function
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}