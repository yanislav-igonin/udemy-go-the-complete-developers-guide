package main

import "fmt"

// Create a new type of 'deck'
// which is slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
