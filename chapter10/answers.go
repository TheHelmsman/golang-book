/*

	1. How do you specify the direction of a channel type?
		
		to send:
			func pinger(c chan<- string)

		to receive:
			func printer(c <-chan string)

	3. What is a buffered channel? How would you create one with a capacity of 20?

		c := make(chan int, 20)
*/
