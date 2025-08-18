package main

import (
	"fmt"
	"packagesapp/myutils"
)

/*
	Packaging serves 3 purposes:
		
	1. It reduces the chance of having overlapping names. 
		This keeps our function names short and succinct

	2. It organizes code so that its easier to find code you want to reuse.

	3. It speeds up the compiler by only requiring recompilation of smaller 
	chunks of a program. Although we use the package fmt, we don't have to 
	recompile it every time we change our program.

	To run code:

		go mod init packagesapp // <- to create `go.mod` file
		go run packages.go //	<- to run application
*/

func main() {
	xs := []float64{1,2,3,4}
	fmt.Println("Source data: ", xs, "\n")
	
	avg := myutils.Average(xs)
	fmt.Println("Avg: ", avg)
	
	min := myutils.Min(xs)
	fmt.Println("Min: ", min)
	
	max := myutils.Max(xs)
	fmt.Println("Max: ", max)
}