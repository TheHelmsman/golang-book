package main

import "fmt"

/*
	Go has a special statement called defer which schedules 
	a function call to be run after the function completes

*/
func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()	//	prints 1st followed by 2nd
}

/*
	defer is often used when resources need to be freed in some way. 
		f, _ := os.Open(filename)
		defer f.Close()

	This has 3 advantages: 
		1) it keeps our `Close` call near our `Open` call so its easier to understand, 
		2) if our function had multiple return statements (perhaps one in an `if` 
		and one in an `else`) `Close` will happen before both of them
		3) deferred functions are run even if a run-time panic occurs.
*/

