package main

import "fmt"

/*
	It is possible to create functions inside of functions:
*/

func main() {
	//	add is a local variable that has the type func(int, int) int
	add := func(x, y int) int {
		return x + y 
	}
	
	fmt.Println(add(1,1))
}

/*
	A function like this together with the non-local variables it references 
	is known as a closure. 
	
	In this case increment and the variable x form the closure:
*/
func main() {
	x := 0
	increment := func() int {
		x++
		return x 
	}
	fmt.Println(increment())
	fmt.Println(increment())
}