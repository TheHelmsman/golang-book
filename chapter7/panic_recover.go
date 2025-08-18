package main

import "fmt"

/*
	We can handle a run-time panic with the built-in recover function. 
	`recover` stops the panic and returns the value that was passed to 
	the call to `panic`.

	wrong implementation:
		func main() {
			panic("PANIC")
			str := recover()
			fmt.Println(str)
		}
	
		But the call to `recover` will never happen in this case because 
		the call to `panic` immediately stops execution of the function. 
		Instead we have to pair it with `defer`
*/

//	correct implementation:
func main() {
	//	wrap recover with defer function:
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}