package main

import (
	"fmt"
	"sort"
)

/*
	The sort package contains functions for sorting arbitrary data. 
	There are several predefined sorting functions (for slices of ints and floats)

	The `Sort` function in sort takes a `sort.Interface` and sorts it. 
	The `sort.Interface` requires 3 methods: `Len`, `Less` and `Swap`.

	To define our own sort we create a new type (`ByName`) 
	and make it equivalent to a slice of what we want to sort. 
	We then define the 3 methods.
*/

type Person struct {
  Name string
	Age int 
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
	}
	fmt.Println("ByName:")
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	
	fmt.Println("ByAge:")
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

