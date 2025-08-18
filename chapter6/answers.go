package main

import (
	"fmt"
	"math"
)

func main() {
	//	How do you access the 4th element of an array or slice?
	foo := [5]int{1,2,3,4,5}
	fmt.Println("4th element: ", foo[3])

	//	What is the length of a slice created using: 
	// 	`make([]int, 3, 9)`
	x := make([]int, 3, 9)
	fmt.Println("len: ", len(x))	// result: 3

	//	what would y[2:5] give you?
	y := [6]string{"a","b","c","d","e","f"}
	fmt.Println(y[2:5])	// result c,d,e,f

	//	smallest number in the given list:
	z := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
	}

	var smallest float64 = math.Inf(1)
	for _, value := range z {
		// fmt.Println("i: ", i, "value: ", value)
		if smallest > float64(value) {
			smallest = float64(value)
			fmt.Println("updating smallest to: ", smallest)
		}
	}

	fmt.Println("smallest final result: ", smallest)
}