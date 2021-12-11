package main

import "fmt"

func printHelloTo(firstName string, lastName string) {
	fmt.Print("Hello ", firstName, " ", lastName, ".")
}

func main() {
	printHelloTo("Syukri", "Hadi")
	fmt.Println()
	printHelloTo("Mutia", "Wardani")
}
