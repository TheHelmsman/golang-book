package main

import (
	"fmt"
	"flag"
	"math/rand"
)

/*
	When we invoke a command on the terminal it's possi-ble to pass that command arguments. 
	We've seen this with the `go` command:

		go run myfile.go

	run and `myfile.go` are arguments. 
	
	We can also pass flags to a command:

		go run -v myfile.go
	
	The flag package allows us to parse arguments and flags sent to our program. 
	Here's an example program that generates a number between 0 and 6. 
	
	We can change the max value by sending a flag (`-max=100`) to the program:
*/

func main() {
	// Define flag
	maxp := flag.Int("max", 6, "the max value")
	fmt.Println("Set flag, default value: ", *maxp)

	// Parse the command-line arguments.
	flag.Parse()
	fmt.Println("Get cli param, update flag: ", *maxp)

	// Generate a number between 0 and max
	fmt.Println("Generate a number between 0 and max:", rand.Intn(*maxp))
}

/*
	go doc flag.Int
package flag // import "flag"

func Int(name string, value int, usage string) *int
    Int defines an int flag with specified name, default value, and usage
    string. The return value is the address of an int variable that stores the
    value of the flag.

	How flag.Int works:
	
	Definition:
		You call flag.Int("flagname", defaultValue, "usage string").
			"flagname": The name of the flag as it will appear on the command line (e.g., -port).
			defaultValue: The integer value the flag will have if it's not provided on the command line.
			"usage string": A brief explanation of what the flag does, displayed when `flag.Usage()` or 
			`flag.PrintDefaults()` is called.

		Return Value:
			`flag.Int` returns a pointer to an `int` variable. This variable will store the value of the flag 
			after `flag.Parse()` is called.

		Parsing:
			You must call `flag.Parse()` after defining all your flags. This function parses the 
			command-line arguments and populates the flag variables with their respective values.

		Accessing the Value:
			You can then dereference the pointer returned by `flag.Int` to access the integer value 
			provided by the user or the default value if no flag was provided.
*/