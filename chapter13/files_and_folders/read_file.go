package main

import (
	"fmt"
	"io/ioutil"
)

/*
  Reading files is very common, so there's a shorter way to do this:
*/

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	
	if err != nil {
		return 
	}
	
	str := string(bs)
	fmt.Println(str)
}

/*
go doc ioutil.ReadFile
package ioutil // import "io/ioutil"

func ReadFile(filename string) ([]byte, error)
	ReadFile reads the file named by filename and returns the contents.
	A successful call returns err == nil, not err == EOF. Because ReadFile
	reads the whole file, it does not treat an EOF from Read as an error to be
	reported.

	Deprecated: As of Go 1.16, this function simply calls os.ReadFile.
*/