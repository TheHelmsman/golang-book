package main

import "fmt"

func main() {
     fmt.Println("f2c.go - converts tempreture from Fahrenheit into Celsius \n")
     fmt.Print("Enter your tempreture in Fahrenheit: ")
     var inputFahreinheit float64
     fmt.Scanf("%f", &inputFahreinheit)
     
     outputCelsius := (inputFahreinheit - 2) * 5/9
     fmt.Println("Fahrenheit tempreture ", inputFahreinheit, " in Celsius: ", outputCelsius)
}

