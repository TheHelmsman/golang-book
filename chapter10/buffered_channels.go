package main

import (
	"fmt"
	"time"
)

/*
	It's also possible to pass a second parameter to the make function when creating a channel:

		c := make(chan int, 1)

	This creates a buffered channel with a capacity of 1. 
	
	Normally channels are synchronous; 
	both sides of the channel will wait until the other side is ready. 
	
	A buffered channel is asynchronous; 
	sending or receiving a message will not wait unless the channel is already full.

	Decoupling:
		Buffered channels allow for a degree of decoupling between sending and receiving goroutines. 
		A sender can continue to send values as long as the buffer is not full, without waiting for 
		a receiver to be ready. Similarly, a receiver can retrieve values from the buffer without 
		the sender being actively available.

	Use cases:
		They are useful in scenarios where you want to handle bursts of data, create producer-consumer 
		patterns where the producer might temporarily outpace the consumer, or to implement patterns 
		like a semaphore.

	Capacity considerations:
		The choice of buffer size is crucial. A buffer size of 0 results in an unbuffered channel. 
		A buffer size of 1 is often used for signaling or to allow a sender to proceed without 
		waiting for an immediate receiver. Larger buffer sizes can be used when a significant 
		amount of temporary storage is needed between goroutines.
*/

func pinger(c chan int) {
	for i := 0; ; i++ {
		c <- 1
	}
}

func printer(c chan int) {
	for {
			fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("This program will print 1 forever (hit ”Enter” to stop it).")
	c := make(chan int, 2)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)	// <- wait for "Enter"
}