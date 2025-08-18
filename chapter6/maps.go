package main

import "fmt"

/*
	A map is an unordered collection of key-value pairs. 
	Also known as an associative array, a hash table or 
	a dictionary, maps are used to look up a value by 
	its associated key.
*/

func main() {
	//	The map type is represented by the keyword map, 
	// 	followed by the key type in brackets and finally 
	// 	the value type: 
	// `map[<key type>]<value type>`
	// var x map[string]int

	//	maps have to be initialized before they can be used
	//	var x map[string]int	//	<- leads to runtime error

	//	proper definition
	// x := make(map[string]int)	// <- no error, make used
	// x["key"] = 10
	// x["foo"] = 20
	// x["bar"] = 20
	// fmt.Println(x["key"])
	// fmt.Println(x["foo"])
	// fmt.Println(x)

	//	create map with int type:
	// x := make(map[int]int)
	// fmt.Println("Map initialized, len:", len(x))	// prints 0
	
	// x[1] = 10
	// fmt.Println(x[1], "assigned")
	// fmt.Println("len: ", len(x), "after assignment")	// prints 1 after assignment
	
	// delete(x, 1)
	// fmt.Println("len: ", len(x), "after deletion")	// prints again after deletion

	//	a very common way of using maps: as a lookup table or a dictionary:
	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["He"] = "Helium"
	// elements["Li"] = "Lithium"
	// elements["Be"] = "Beryllium"
	// elements["B"] = "Boron"
	// elements["C"] = "Carbon"
	// elements["N"] = "Nitrogen"
	// elements["O"] = "Oxygen"
	// elements["F"] = "Fluorine"
	// elements["Ne"] = "Neon"

	//	short form of maps creation:
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
	
	fmt.Println(elements["Li"])
	//	Lithium
	fmt.Println("[\"Un\"] element: ", elements["Un"])	// nothing returned
	//	["Un"] element:
	
	//	Accessing an element of a map can return two values instead of just one. 
	//	The first value is the result of the lookup, the second tells us whether 
	//	or not the lookup was successful.
	name, ok := elements["Un"]
	fmt.Println("[\"Un\"] element: ", name, ok)				// false returned
	//	["Un"] element:  false

	if name, ok := elements["Un"]; ok {
		fmt.Println("Print result with lookup: ", name, ok)
		//	no output, condition failed
	}

	shortElements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}
	if el, ok := shortElements["Li"]; ok {
		fmt.Println("short elements: ", el["name"], el["state"])
		//	short elements:  Lithium solid
	}
}