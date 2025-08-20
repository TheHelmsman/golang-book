package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	Often we want to recursively walk a folder
	 (read the folder's contents, all the sub-folders, all the sub-sub- folders, ...). 
	 To make this easier there's a `Walk` function provided in the `path/filepath` package
*/

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil 
	})
}
