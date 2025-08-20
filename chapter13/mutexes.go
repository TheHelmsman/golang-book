package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	The preferred way to handle concurrency and synchronization in Go is 
	through goroutines and channels as discussed in chapter 10. 
	
	However Go does provide more traditional multithreading routines in the 
	`sync` and `sync/atomic` packages.

	Mutexes:

	A mutex (mutal exclusive lock) locks a section of code to a single thread 
	at a time and is used to protect shared resources from non-atomic operations. 
	
	Here is an example of a mutex:
*/

func main() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i) 
	}
	var input string
	fmt.Scanln(&input)
}