package main

import "fmt"

func main() {
	// variable declaration with type
	var name string

	name = "Syukri Hadi"
	fmt.Println(name)

	name = "Mutia Kusuma Wardani"
	fmt.Println(name)

	// golang will auto detect data type if the value assigned when the variable being declared
	var age = 19
	fmt.Println(age)

	var usia int8 = 22
	fmt.Println(usia)

	// another way to declare a variable using syntax :=
	friendName := "Kamil"
	fmt.Println(friendName)

	friendName = "Pasaribu"
	fmt.Println(friendName)

	// multiple variable declaration
	var (
		firstName = "Syukri Hadi"
		lastName  = "Kamil"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
