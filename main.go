package main

func main() {
	//str := "learn go"
	//temp := &str
	//fmt.Printf("%d %s %s", &str, *temp, str)
	//cards := []string{"temp"}

	//for index loop
	//for i := 0; i < 5; i++ {
	//	card := newCard(i)
	//	fmt.Println(i, card)
	//}

	//for range loop
	//for i, c := range cards {
	//	fmt.Println(i, c)
	//}

	//fmt.Println(cards)

	fruits := getFruits()
	//someOfFruits := fruits[0:2] // new slice of fruits from index 0 (include) to index 2 (not include) => [apple,banana]
	//someOfFruits := fruits[:2] it's also starting from index 0
	//someOfFruits := fruits[2:] it's starts from index 2 to end index of slice => [orange kiwi pomegranate strawberry]

	//fmt.Println(someOfFruits)

	bag1, _ := fruitsBag(fruits, 3)

	//bag1.print()
	//convert to byte slice
	//sbag1 := strings.Join(bag1, ",")
	//bbag1 := []byte(sbag1)
	//fmt.Println(bbag1)

	//bag2.print()

	//ignore returning value
	//_ = bag1.saveToFile("bag1")

	//error handling when byte slice writing to file
	//err := bag2.saveToFile("bag2")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}

	//error handling when file content reading to byte slice
	//bs, err := bag1.readFromFile("bag1")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	//exit program with error code 1
	//	os.Exit(1)
	//}
	//
	//bag1FromBs := strings.Split(string(bs), ",")
	//
	//fmt.Println(bag1FromBs)

	bag1.print()
	bag1.shuffle()
	bag1.print()
}
