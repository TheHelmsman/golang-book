package main

import (
	"fmt"
	"os"
)

/*
	To open a file in Go use the Open function from the `os` package.
*/

func main() {
	file, err := os.Open("input_output.go")
	
	if err != nil {
		// handle the error here
		return
	}
	
  //  We use defer file.Close() right after opening the file to make sure 
  // the file is closed as soon as the function completes.
	defer file.Close()
	
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	
	if err != nil {
		return
	}
	
	str := string(bs)
	fmt.Println(str)
}

