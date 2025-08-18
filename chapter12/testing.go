package main

import (
	"fmt"
	"testing_pckg/math"
)

func main() {
	xs := []float64{}
	fmt.Println("Source data: ", xs, "\n")
	
	avg := math.Average(xs)
	fmt.Println("Avg: ", avg)
}