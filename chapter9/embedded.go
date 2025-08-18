package main

import ("fmt")

/*
	A struct's fields usually represent the `has-a` relationship,
	while embedded types represents the `is-a` relationship
*/

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
/*
	Creating struct:

		type Android struct {
			Person Person
			Model string
		}

	This would work, 
	but we would rather say an Android `is` a Person, 
	rather than an Android `has` a Person.

	Therefore we use embedded type:
*/

type Android struct {
	Person	//	<- We use the type (Person) and don't give it a name.
	Model string
}

func main() {
	// a := new(Android)
	// a.Person.Talk()
	
	p := Person{"Michael"}
	p.Talk()
	
	//	But we can also call any Person methods directly on the Android:
	// a := Android{Person{"Android"}, "x"}
	a := new(Android)
	a.Name = "Android"
	a.Talk()
}