package main

import "fmt"

type deck []string

//receiver function
//By creating a new type with a function that has a receiver, we are adding a 'method' to any value of that type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

	// _ if you want to not use index use underscore
	//for _, card := range d {
	//	fmt.Println(card)
	//}
}
