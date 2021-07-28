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

	alexPointer := &alex
	alexPointer.updateName("Loh")

	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
