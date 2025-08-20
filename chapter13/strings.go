package main

import (
  "fmt"
	"strings"
)

/*
	Go includes a large number of functions to work with strings in the strings package

	Sometimes we need to work with strings as binary data. 
	To convert a string to a slice of bytes (and vice- versa) do this:

	arr := []byte("test")
	str := string([]byte{'t','e','s','t'})
*/
func spacer() {
	fmt.Println("\n")
}

func main() {
	fmt.Println(
		// true
		"test Contains es: ", strings.Contains("test", "es"), "\n",
		// 2
		"test count t: ", strings.Count("test", "t"), "\n",
		// true
		"test HasPrefix te: ", strings.HasPrefix("test", "te"), "\n",
		// true
		"test HasSuffix st: ", strings.HasSuffix("test", "st"), "\n",
		// 1
		"test Index e: ", strings.Index("test", "e"), "\n",
		// "a-b"
		"Join a b: ", strings.Join([]string{"a","b"}, "-"), "\n",
		// == "aaaaa"
		"Repeat a: ", strings.Repeat("a", 5), "\n",
		// "bbaa"
		"Replace aaa, a, b: ", strings.Replace("aaaa", "a", "b", 2), "\n",
		// []string{"a","b","c","d","e"}
		"Split a-b-c-d-e: ", strings.Split("a-b-c-d-e", "-"), "\n",
		// "test"
		"ToLower: ", strings.ToLower("TEST"), "\n",
		// "TEST"
		"ToUpper: ", strings.ToUpper("test"),
	)
	
	spacer()
	arr := []byte("test")
	fmt.Println("convert to array, source string: test")
	fmt.Println("array (bytes): ", arr)
	
	spacer()
	str := string([]byte{'t','e','s','t'})
	fmt.Println("convert array (bytes) to string, source array: ", []byte{'t','e','s','t'})
	fmt.Println("string from array: ", str)
}