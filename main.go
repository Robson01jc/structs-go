package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

// This function takes a pointer to a person as receiver.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// A struct is a collection of properties that are related together.
// We use the 'type' keyword to define custom types.
// When you are defining multiline structures, every single line must end with a comma.
// RAM is like a bunch of slots/boxes, each one of which has a discrete address.
// Go is a 'pass by value' language.

// &variable -> get the memory address of the value this variable is pointing at.
// *pointer  -> get the value this memory address is pointing at.
