package main

import "fmt"

/*
	A switch statement starts with the keyword 
	switch followed by an expression (in this 
	case i) and then a se- ries of cases. 
	The value of the expression is compared to 
	the expression following each case keyword. 
	If they are equivalent then the statement(s) 
	following the : is executed.

*/
func main() {
	for i := 1; i <= 6; i++ {
		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown number")
		}
	}
}