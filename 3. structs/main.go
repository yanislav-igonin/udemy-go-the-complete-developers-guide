package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // shortcut
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			zip:   900010,
			email: "alex@alex.com",
		},
	}

	alex.updateName("Loh")

	alex.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
