package main

import "fmt"

/*
	First we need a way of determining whether 
	or not a number is even or odd. An easy way 
	to tell is to divide the number by 2. 
	If you have nothing left over then the number 
	is even, otherwise it's odd. So how do we 
	find the remainder after division in Go? 
	We use the % operator.
		1 % 2 equals 1,
		2 % 2 equals 0,
		3 % 2 equals 1
	and so on.
*/

func main() {
	//	loop
	for i := 1; i <= 10; i++ {
		//	check
		if i % 2 == 0 {
			//	even
			fmt.Println(i, "even")
		} else {
			//	odd
			fmt.Println(i, "odd")
		}
	}
}

/*
	If statements also have an optional else part. 
	If the condition evaluates to true then the 
	block after the condition is run, otherwise 
	either the block is skipped or if the else 
	block is present that block is run.

	if i % 2 == 0 {
    // divisible by 2
	} else if i % 3 == 0 {
		// divisible by 3
	} else if i % 4 == 0 {
		// divisible by 4
	}
*/