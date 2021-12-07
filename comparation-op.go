package main

import "fmt"

func main() {
	var name1 string = "Syukri"
	var name2 = "Syukri"
	name3 := "Mutia"

	var result1 bool = name1 == name2
	var result2 = name1 == name3

	fmt.Println(result1)
	fmt.Println(result2)

	// number comparation

	var value1 int16 = 100
	var value2 int16 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
