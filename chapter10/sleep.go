package main

import (
	"fmt"
	"time"
)

/*
	2. Write your own `Sleep` function using `time.After`.

	func MySleep(d time.Duration):
		This defines the custom Sleep function that takes a time.Duration as an argument, similar to time.Sleep.
	
		<-time.After(d):
		This is the core of the implementation.
			
			time.After(d) creates a new Timer and returns a <-chan time.
			Time (a receive-only channel of time.Time). This channel will receive the current 
			time value after the duration d has passed.
			
			The <- operator attempts to receive a value from this channel. 
			Since the channel will only send a value after the specified duration, the goroutine 
			calling MySleep will block at this line until the value is sent, effectively pausing 
			its execution for the desired duration.
*/

// MySleep pauses the current goroutine for at least the duration d.
// It uses time.After and a channel for implementation.
func MySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Starting sleep...")
	MySleep(2 * time.Second) // Pause for 2 seconds
	fmt.Println("Sleep finished!")

	fmt.Println("Starting another sleep...")
	MySleep(500 * time.Millisecond) // Pause for 500 milliseconds
	fmt.Println("Another sleep finished!")
}