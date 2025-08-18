package main

import "fmt"

/*
	types:
		
	- booleans
	/	a special 1 bit integer type used to represent true and false /

	Three logical operators are used with boolean values:
		&& - and
		|| - or
		! - not

	Truth table:
		true && true		- true
		true && false		- false
		false && true 	- false
		false && false	- false
		true || true 		- true 
		true || false 	- true 
		false || true 	- true
		false || false	- false
		!true						- false
		!false					- true
*/
func main() {
	fmt.Println(true && true)
	//	output: true
	
	fmt.Println(true && false)
	//	output: false

	fmt.Println(true || true)
	//	output: true
	
	fmt.Println(true || false)
	//	output: true
	
	fmt.Println(!true)
	//	output: false
}