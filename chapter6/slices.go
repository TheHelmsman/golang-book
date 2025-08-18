package main

import "fmt"

/*
	A slice is a segment of an array. 
	Like arrays slices are indexable and have a length. 
	Unlike arrays this length is allowed to change.

	Go includes two built-in functions to assist with slices: 
	- `append` 
	- `copy`
*/

func main() {
	//	slice
	// var x []float64

	//	create slice with build in "make" function
	//	the following line creates a slice that is 
	// 	associated with an underlying float64 array 
	// 	of length 5
	// x := make([]float64, 5)

	//	create slice with underlying array
	//	10 represents the capacity of the underlying 
	// 	array which the slice points to
	// x := make([]float64, 5, 10)
	//	[x,x,x,x,x,_,_,_,_,_]

	//	create slice with [low:high] syntax
	// arr := []float64{1,2,3,4,5}
	// x := arr[0:5]
	//	arr[0:5] returns 1,2,3,4,5
	//	arr[1:4] returns 2,3,4
	//	arr[:5] is the same as arr[0:5]
	//	arr[:] is the same as arr[0:len(arr)]

	//	append 
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Append result: ", slice1, slice2)
	//	Append result:  [1 2 3] [1 2 3 4 5]

	//	copy slice3 into slice4 
	//	slice4 has room only for 2 elements
	//	so it remains to have 2 elements after copying
	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Copy result: ", slice3, slice4)
	//	Copy result:  [1 2 3] [1 2]
}