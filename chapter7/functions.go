package main

import "fmt"

/*
	A function is an independent section of code 
	that maps zero or more input parameters to 
	zero or more output parameters.

	The parameters (inputs) of the function are 
	defined like this: `name type, name type, ...`

	Collectively the parameters and the return type 
	are known as the `function's signature`.

	Functions don't have access to anything in the 
	calling function.
*/


/* 
	Find average in array
	@param data	[]float64 - input area
 	returns float64 - average calculation

	alternative: func average(data []float64) (result float64) {}
*/
func average(data []float64) float64 {
	// panic("Not Implemented")

	total := 0.0
	for _, v := range data {
		total += v
	}
	return total / float64(len(data))
}

func main() {
	xs := []float64{ 98,93,77,82,83 }
	fmt.Println(average(xs))
}

//	we can also name a return type:
// func f2() (r int) { 
// 	r = 1
// 	return 
// }