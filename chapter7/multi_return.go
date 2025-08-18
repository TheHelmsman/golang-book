package main

import "fmt"

/*
	Go is also capable of returning multiple values from a function
*/

func f() (int, int) {
	//	return multiple values separated by ,
	return 5, 6
}

func main() {
	//	assign returns:
	x, y := f()
	fmt.Println("x:",x,"y:",y)
}

//	Multiple values are often used to 
//	1.	return an error value along with the result (x, err := f()), 
//	2.	or a boolean to indicate success (x, ok := f()).
