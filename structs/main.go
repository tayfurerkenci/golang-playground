package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	tyfr := person{
		firstName: "Tayfur",
		lastName:  "Erkenci",
		contactInfo: contactInfo{
			email:   "terkenci@gmail.com",
			zipCode: 12345,
		},
	}

	// turn address into value with *address
	// turn value into address with &value
	tyfr.updateName("Taylor")

	tyfr.print()
}

// *person is a type description, it means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson is an operator, it means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
