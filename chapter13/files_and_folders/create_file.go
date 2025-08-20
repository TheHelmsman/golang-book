package main

import ( 
	"os"
)

/*
	Here is how we can create a file:
*/

func main() {
	file, err := os.Create("test.txt")
	
	if err != nil {
		// handle the error here
		return 
	}
	
	defer file.Close()
	
	file.WriteString("test")
}