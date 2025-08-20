package main

import (
	"fmt"
	"hash/crc32"
)

/*
	A hash function takes a set of data and reduces it to a smaller fixed size. 
	Hashes are frequently used in programming for everything from looking up data 
	to easily detecting changes. 
	
	Hash functions in Go are broken into two categories: `cryptographic` and `non-cryptographic`.
	The `non-cryptographic` hash functions can be found underneath the `hash` package and include 
	
		adler32, crc32, crc64 and fnv. 
	
	Here's an example using `crc32`:
*/

func main() {
	//	A common use for crc32 is to compare two files. 
	// If the Sum32 value for both files is the same, it's highly likely (though not 100% certain) 
	// that the files are the same. 
	// If the values are different then the files are definitely not the same
	h := crc32.NewIEEE()
	
	// The `crc32` hash object implements the `Writer` interface, 
	// so we can write bytes to it like any other `Writer`.
	h.Write([]byte("test"))
	
	//	Once we've written everything we want we call `Sum32()` to return a `uint32`.
	v := h.Sum32()
	
	fmt.Println("CRC of \"test\" is:", v)
}