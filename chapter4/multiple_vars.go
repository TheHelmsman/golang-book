package main

import "fmt"

/*
  Go also has another shorthand when you need to de- fine multiple variables:
  
  var ( 
    a = 5
    b = 10
    c = 15 
  )

*/

func main() {
  
  fmt.Print("Enter a number: ")
  var input float64
  
  fmt.Scanf("%f", &input)
  
  output := input * 2
  fmt.Println("your input:", input, "x2 =", output)
}

/*

$go doc fmt.Scanf 

Output:
package fmt // import "fmt"

func Scanf(format string, a ...any) (n int, err error)
    Scanf scans text read from standard input, storing successive
    space-separated values into successive arguments as determined by the
    format. It returns the number of items successfully scanned. If that is less
    than the number of arguments, err will report why. Newlines in the input
    must match newlines in the format. The one exception: the verb %c always
    scans the next rune in the input, even if it is a space (or tab etc.) or

*/