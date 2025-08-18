package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
	Making progress on more than one task simultaneously is known as concurrency. 
	Go has rich support for concurrency using goroutines and channels.

	A goroutine is a function that is capable of running concurrently with other functions. 
	To create a goroutine we use the keyword `go` followed by a function invocation:

	go f(0)
*/

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	} 
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i) 
	}
	// go f(0)	//	<- concurrent routine
	
	var input string
	//	without `Scanln` the program would exit before being given 
	//  the opportunity to print all the numbers.
	fmt.Scanln(&input) 
}