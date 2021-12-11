package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Syukri"
	middleName = "Hadi"
	lastName = "Kamil"

	return // no need to specify the returned variables
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
