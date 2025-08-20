package main

import (
	"fmt"
	"container/list"
)

/*
	The `container/list package` implements a doubly-linked list. 
	A linked list is a type of data structure that looks like this:

	Front - Node 1 -> Node 2 -> Node 3 - Back
						1					2					3

	Each node of the list contains a value (1, 2, or 3 in this case) 
	and a pointer to the next node.

	Since this is a doubly-linked list each node will also have pointers to the previous node.

	The zero value for a List is an empty list (a *List can also be created using `list.New`). 
	Values are appended to the list using `PushBack`. 
*/

func main() {
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	} 
}