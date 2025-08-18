package main

import (
	"fmt"
	"time"
)

/*
	Making progress on more than one task simultaneously is known as concurrency. 
	Go has rich support for concurrency using goroutines and channels.

	Channels provide a way for two goroutines to communicate with one another 
	and synchronize their execution.

	A channel type is represented with the keyword `chan` followed by the 
	type of the things that are passed on the channel (for instance `strings`).

	The `<-`` (left arrow) operator is used to send and receive messages on the channel.

	Using a channel like this synchronizes the two goroutines.

	When pinger attempts to send a message on the channel it will wait until printer 
	is ready to receive the message (this is known as blocking).

	We can specify a direction on a channel type thus re- stricting it to either sending or receiving.

	to send:
		func pinger(c chan<- string)

	to receive:
		func printer(c <-chan string)
	
	A channel that doesn't have these restrictions is known as bi-directional.
*/

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" 
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		//	fmt.Println(<-c) //	in this case msg := <- c is not needed
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("This program will print “ping” forever (hit ”Enter” to stop it).")
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)	// <- wait for "Enter"
}