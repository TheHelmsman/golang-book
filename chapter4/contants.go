package main

import "fmt"

/*
	Constants are basi- cally variables whose values cannot be changed later.
	/ created using the `const` keyword /

*/

func main() {
	const x string = "Hello World"
	//	x = "Some other string" -> Results in a compile-time error:
	//	.\main.go:7: cannot assign to x
	fmt.Println(x)
}

