package main

import "fmt"
import "os"

//	this is a comment)

/*
	this is also a comment)
*/

func main() {
	
	//	print args amount, 1 - no args, >=2 - args exists
	// fmt.Println(len(os.Args))
	
	var bytes int	// declare var with int type
	
	if(len(os.Args) >= 2) {
		//	fmt.Println function returns 2 values: int, error, we grab first and discard second
		bytes, _ = fmt.Println("Hello, my name is", os.Args[1])	//	save amount of bytes printed
	} else {
		bytes, _ = fmt.Println("Hello, world")	//	save amount of bytes printed
	}
	
	fmt.Println(bytes, "bytes written")	//	print amount of bytes written by fmt.Println
	os.Exit(0)	// immediate program termination: code zero indicates success, non-zero an error
}


/*

NOTE: os.Args provides access to raw command-line arguments. 
Note that the first value in this slice is the path to the program, 
and os.Args[1:] holds the arguments to the program.

argsWithProg := os.Args
argsWithoutProg := os.Args[1:]
arg := os.Args[3]
fmt.Println(argsWithProg)
fmt.Println(argsWithoutProg)
fmt.Println(arg)

Output:
[/tmp/go-build162373819/command-line-arguments/_obj/exe/modbus 1 2 3 4 5]
[1 2 3 4 5]
3

NOTE: go doc provides access to go documentation
for example, to get help on os.Exit function write: 

$go doc os.Exit

>Output:
package os // import "os"

func Exit(code int)
    Exit causes the current program to exit with the given status code.
    Conventionally, code zero indicates success, non-zero an error. The program
    terminates immediately; deferred functions are not run.

    For portability, the status code should be in the range [0, 125].
*/