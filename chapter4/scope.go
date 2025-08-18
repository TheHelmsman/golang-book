package main

import "fmt"

var x string = "Hello World"

func main() {
    fmt.Println("x: ", x)
	f()
}

func f() {
    fmt.Println("Another execution, x:", x)
}

/*

func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
func f() {
    fmt.Println(x)
}

If you run this program you should see an error:
.\main.go:11: undefined: x

*/