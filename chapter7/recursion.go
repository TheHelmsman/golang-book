package main

import "fmt"

/*
	A function is able to call itself
*/

func factorial(x uint) uint {
	if x == 0 {
		return 1 
	}
	return x * factorial(x - 1)
}

func main() {
	fmt.Println("factorial(2): ", factorial(2))
}

// factorial(2):
// 	Is x == 0? No. (x is 2)
// 	Find the factorial of x – 1
// 	Is x == 0? No. (x is 1)
// 	Find the factorial of x – 1
// 	Is x == 0? Yes, return 1. 
// 	return 1*1
// 	return 2*1
