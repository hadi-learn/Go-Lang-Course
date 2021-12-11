package main

import "fmt"

func getFullName() (string, string, string) {
	return "Syukri", "Hadi", "Kamil"
}

func main() {
	firstName, _, lastName := getFullName() // using _ to ignore a return value
	fmt.Println(firstName, lastName)
}
