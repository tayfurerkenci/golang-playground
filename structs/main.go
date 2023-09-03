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

	tyfr.updateName("Taylor")

	tyfr.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
