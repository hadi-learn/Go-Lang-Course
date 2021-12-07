package main

import "fmt"

func main() {
	// constant declaration
	const firstName string = "Syukri"
	const lastName = "Hadi Kamil"
	const age = 35

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// multiple constant declaration
	const (
		wifeName = "Mutia Kusuma Wardani"
		bikeName = "Vario 125 cc"
	)

	fmt.Println(wifeName)
	fmt.Println(bikeName)
}
