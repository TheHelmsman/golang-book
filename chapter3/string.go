package main

import "fmt"

/*
	types:
		
	- string
	/ a sequence of char- acters with a definite length used to represent text. /

	can use "" or ``
	you can place \n and \t inside strings with ``
*/

func main() {
	//	get length of a string, space considered as a char
	fmt.Println("length: ", len("Hello World"))
	//	output: 11

	//	access individual char in string, strings indexed starting from 0
	//	101 instead of e is because character is represented by a byte
	fmt.Println("Hello World"[1])
	//	output: 101

	//	concatenate 2 strings together, note space after Hello
	fmt.Println("Hello " + "World")
	//	output: Hello World
}