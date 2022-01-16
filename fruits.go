package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type fruits []string

func getFruits() fruits {
	return fruits{"apple", "banana", "orange", "kiwi", "pomegranate", "strawberry", "pear", "peach", "plum", "tangerine"}
}

//that is a receiver function. "fruits" type of string slice and every "fruits" type include "print" via that receiver function
func (f fruits) print() {
	fmt.Println(f)
}

func (f fruits) toString() string {
	return strings.Join(f, ",")
}

func (f fruits) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(f.toString()), 0666)
}

func (f fruits) readFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func fruitsBag(f fruits, size int) (fruits, fruits) {
	return f[:size], f[size:]
}

func (f fruits) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range f {

		// alwaws gives same result
		//newPosition := rand.Intn(len(f) - 1)

		newPosition := r.Intn(len(f) - 1)

		//swap the elements
		f[i], f[newPosition] = f[newPosition], f[i]

		//long way 4 swap
		//temp := f[i]
		//f1 := f[newPosition]
		//f[i] = f1
		//f[newPosition] = temp

	}
}
