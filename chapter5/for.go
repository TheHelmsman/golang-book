package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//	could also be written as:
	// i := 1
	// for i <= 10 { fmt.Println(i); i = i + 1 }
	//	also possible:
	// for i <= 10 { fmt.Println(i); i += 1 }
	// for i <= 10 { fmt.Println(i); i++ }
}

// func main() {
// 	fmt.Println(`1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10`)
// }

// func main() {
// 	fmt.Println(1)
// 	fmt.Println(2)
// 	fmt.Println(3)
// 	fmt.Println(4)
// 	fmt.Println(5)
// 	fmt.Println(6)
// 	fmt.Println(7)
// 	fmt.Println(8)
// 	fmt.Println(9)
// 	fmt.Println(10)
// }


