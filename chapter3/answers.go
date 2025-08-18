package main

import "fmt"
import "math"

/*
	We know that (in base 10) the largest 1 digit number is 9 
	and the largest 2 digit number is 99. 
	
	Given that in binary the largest 2 digit number is 11 (3), 
	the largest 3 digit number is 111 (7) 
	and the largest 4 digit number is 1111 (15) 
	what's the largest 8 digit number? (hint: 10^1-1 = 9 and 10^2-1 = 99)

	answer: 
		in base 10: 10^8-1 = 99999999
		in base 2: 1*2^7+1*2^6+1*2^5+1*2^4+1*2^3+1*2^2+1*2^1+1*2^0 = 11111111 (255)
*/

func main() {

	//	math.Pow(base, exponent)
	fmt.Println("10^8-1=", math.Pow(10, 8) - 1)

	fmt.Println("321325*424521=", 321325*424521)

	fmt.Println("String's length:", len("String's length:"))

	fmt.Println("Expression value: ", (true && false) || (false && true) || !(false && false))
}