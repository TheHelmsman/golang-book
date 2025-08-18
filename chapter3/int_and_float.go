package main

import "fmt"

/*
    Go is a statically typed programming language. 
    This means that variables always have a specific type 
    and that type cannot change.

    types:
        - integers
        / numbers without decimal component /
        
        Go's integer types are: 
        uint8, uint16, uint32, uint64, int8, int16, int32 and int64.
        
        uint means “un- signed integer” while int means “signed integer”. 
        Un- signed integers only contain positive numbers (or zero).

        - floating point numbers 
        / numbers that contain a decimal component (real numbers) /

        Go has two floating point types: 
        float32 and float64
        and two additional complex numbers types:
        complex64 and complex128.
*/

func main() {
    fmt.Println("1 + 1 =", 1.0 + 1.0) // integers
    //fmt.Println("1 + 1 =", 1.0 + 1.0) // floats
}