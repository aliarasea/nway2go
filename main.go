package main

import "fmt"

func main() {
	str := "learn go"
	temp := &str
	fmt.Printf("%d %s %s", &str, *temp, str)
	cards := []string{"temp"}

	//for index loop
	for i := 0; i < 5; i++ {
		card := newCard(i)
		fmt.Println(i, card)
	}

	//for range loop
	for i, c := range cards {
		fmt.Println(i, c)
	}

	fmt.Println(cards)
}

func newCard(index int) string {
	return fmt.Sprint("Five of diamonds ", index)
}
