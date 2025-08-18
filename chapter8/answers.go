package main

import "fmt"

/*
  1. How do you get the memory address of a variable?

  `new()` or `&`

  2. How do you assign a value to a pointer?

  `*nameOfVar = newValue;`

  3. How do you create a new pointer?

  `newPointer *Type`

  4. What is the value of x after running this program:

    func square(x *float64) {
      *x = *x * *x
    }

    func main() {
      x := 1.5
      square(&x)
      fmt.Println(x)  //  prints 2.25
    }
*/

/*
  5. Write a program that can swap two integers 
  (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
*/

func swap(x *int, y *int) {
  z := *x
  *x = *y
  *y = z
}

func main() {
  x := 1
  y := 2
  swap(&x, &y)
  fmt.Println("x =", x, "y =", y)
}