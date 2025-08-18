package main

import ("fmt"; "math")

/*
	Like a `struct` an `interface` is created using the `type` keyword, 
	followed by a `name` and the keyword `interface`. 
	
	But instead of defining fields, we define a “method set”. 
	A method set is a list of methods that a `type` must have in order to 
	“implement” the interface.

		type Shape interface {
     area() float64
		}

	By itself this wouldn't be particularly useful, 
	but we can use interface types as arguments to functions:

		func totalArea(shapes ...Shape) float64 {}
	
	We would call this function like this:

		fmt.Println(totalArea(&c, &r))
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
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area 
}

//	Interfaces can also be used as fields:
type MultiShape struct {
	shapes []Shape
}

//	We can even turn `MultiShape` itself into a `Shape` by giving it an area method:
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	m := MultiShape{[]Shape{&c, &r}}
	fmt.Println("c area:", c.area())
	fmt.Println("r area:", r.area())
	fmt.Println("total:", totalArea(&c, &r))
	fmt.Println("total:", m.area())
}