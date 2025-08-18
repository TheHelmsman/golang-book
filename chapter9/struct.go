package main

import ("fmt"; "math")

/*
	A struct is a type which contains named fields.

		type Circle struct {
			x float64
			y float64
			r float64 
		}
	
	The `type` keyword introduces a new type. 
	
	It's followed by the name of the type (`Circle`), the keyword `struct` 
	to indicate that we are defining a `struct` type and a list of fields 
	inside of curly braces. 
	
	Each field has a `name` and a `type`.

	We can create an instance of our new Circle type in a variety of ways:
		
		var c Circle

	We can also use the new function:
		
		c := new(Circle)

	This allocates memory for all the fields, sets each of them to their 
	zero value and returns a pointer.

	More often we want to give each of the fields a value. 

		c := Circle{x: 0, y: 0, r: 5}
		c := Circle{0, 0, 5} <- alternative, we can leave off the field names

	We can access fields using the `.` operator
*/

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

/*
	with pointers:

		func circleArea(c *Circle) float64 {
			return math.Pi * c.r*c.r
		}
*/
func main() {

	// var c Circle

	//	alternative
	// c := new(Circle)

	//	alternative
	// c := Circle{x: 0, y: 0, r: 5}

	//	alternative
	// c := Circle{0, 0, 5}

	//	We can access fields using the . operator:
	// fmt.Println(c.x, c.y, c.r)
	// c.x = 10
	// c.y = 5

	//	Let's modify the circleArea function so that it uses a Circle:
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
	//	with pointers
	// fmt.Println(circleArea(&c))
}
   