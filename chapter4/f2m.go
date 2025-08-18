package main

import "fmt"

func main() {
	fmt.Println("f2m.go - converts feets into meters \n")
	const feet float64 = 0.3048
	fmt.Print("Enter you number (in feets): ");
	var feetsInput float64
	fmt.Scanf("%f", &feetsInput)

	result := feetsInput * feet
	fmt.Println("Your feets value: ", feetsInput, " in meters: ", result)
}