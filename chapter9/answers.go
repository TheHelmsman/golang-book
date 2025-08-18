package main

import ("fmt"; "math")

/*
	1. What's the difference between a method and a function?

		In Go, both functions and methods are blocks of code designed to perform specific tasks, but they differ 
		fundamentally in their association with types.

		Functions:
			A function is a standalone block of code that is not associated with any specific data type.
			Functions are declared using the func keyword and can accept parameters and return values.
			They are called directly by their name, and any data they operate on must be explicitly passed as arguments.

			func add(a, b int) int {
				return a + b
			}
	
		Methods:
			A method is a function that is associated with a specific data type, typically a struct.
			Methods are declared with a receiver argument, which specifies the type the method operates on. 
			This receiver can be either a value or a pointer to the type.
			Methods are called on an instance of their associated type, allowing them to access and potentially modify the 
			data within that instance. 

			type Rectangle struct {
				width, height int
			}

			func (r Rectangle) Area() int { // Area is a method with a Rectangle receiver
				return r.width * r.height
			}

		Key Differences Summarized:
			Association:
				Functions are standalone; methods are associated with a specific type (via a receiver).
			Syntax:
				Methods include a receiver argument in their declaration.
			Invocation:
				Functions are called directly; methods are called on an instance of their associated type.
			Data Access:
				Methods can implicitly access the fields of their receiver type, while functions require all 
				data to be explicitly passed as arguments.
		
	2. Why would you use an embedded anonymous field instead of a normal named field?

		In Go, an embedded anonymous field in a struct provides a form of composition that 
		promotes fields and methods of the embedded type directly to the enclosing struct, 
		offering conciseness and promoting code reuse.

	3. Add a new method to the `Shape` interface called `perimeter` which calculates 
	the perimeter of a shape.  Implement the method for `Circle` and `Rectangle`.
*/

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

//	interface
type Shape interface {
	area() float64
	perimeter() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println("c area:", c.area())
	fmt.Println("r area:", r.area())

	fmt.Println("c perimeter:", c.perimeter())
	fmt.Println("r perimeter:", r.perimeter())
}