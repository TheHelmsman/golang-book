package main

import "fmt"

/*
	An array is a numbered sequence of elements 
	of a single type with a fixed length.

	Arrays are indexed starting from 0.
*/

func main() {
	/*
		This program computes the average of a 
		series of test scores. If you run it you 
		should see 86.6.
	*/
	x := [5]float64{ 
		98, 
		93, 
		77, 
		82, 
		83, // <- leave trailing , in the last element
		// 99,
	}
	// var x [5]float64
	
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83

	var total float64 = 0

	/*
		In this for loop _ (or i) represents the current 
		position in the array and value is the same as x[i]. 
		We use the keyword range followed by the name of 
		the variable we want to loop over.
		to avoid i declared and not used, replace i with _
	*/
	for _, value := range x {
		total += value
	}
	//	alternative:
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	//	example of type conversion: 
	// total is a `float64`, while len(x) is `int`
	fmt.Println(total / float64(len(x)))
}