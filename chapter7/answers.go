package main

import (
	"fmt"
	"math"
)
/*
	1. sum
*/
func sum(a []int) int {
	total := 0
	for _, value := range a {
		total += value;
	}
	return total;
}
/*
	2. half
*/
func half( x int ) (int, bool) {
	half := x / 2
	if half % 2 == 1 {
		return half, true
	} else {
		return half, false
	}
}
/*
	3. biggest
*/
func biggest(args ...int) int {
	var biggest float64 = math.Inf(-1)
	for _, value := range args {
		// fmt.Println("i: ", i, "value: ", value)
		if biggest < float64(value) {
			biggest = float64(value)
			// fmt.Println("updating biggest to: ", biggest)
		}
	}
	return int(biggest)
}
/*
	4. makeOddGenerator
		makeOddGenerator returns a function which generates odd numbers. 
		Each time it's called it adds 1 to the local i variable which – unlike 
		normal local variables – persists between calls.
*/
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
/*
	5. Fibbonacci
*/
func fib(x uint) uint {
	if x == 0 {
		return 0 
	}
	if x == 1 {
		return 1 
	}
	return fib(x - 1) + fib(x - 2)
}

func spacer() {
	fmt.Println("\n")
}

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println("sum 1..5 is: ", sum(s))
	spacer()

	var value1, boolean1 = half(1)
	var value2, boolean2 = half(2)
	fmt.Println("half(1): ", "(", value1, ",", boolean1, ")")
	fmt.Println("half(2): ", "(", value2, ",", boolean2, ")")
	spacer()

	arr := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("biggest of ", arr, " is", biggest(arr...))
	spacer()

	nextOdd := makeOddGenerator()
	n 			:= "nextOdd"
	fmt.Println(n, nextOdd()) // 1
	fmt.Println(n, nextOdd()) // 3
	fmt.Println(n, nextOdd()) // 5
	spacer()

	fmt.Println("fib(9)", fib(9))
}
