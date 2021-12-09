package main

import "fmt"

func main() {
	name := "Kamil"

	if name == "Syukri" {
		fmt.Println("Hello", name)
	} else if name == "Hadi" {
		fmt.Println("Hi", name)
	} else {
		fmt.Println("Hi, what's your name?")
	}

	// if statement normal
	var length = len(name)
	if length > 5 {
		fmt.Println("Panjang")
	} else {
		fmt.Println("Pendek")
	}

	// if statement short
	if length2 := len(name); length2 >= 5 {
		fmt.Println("Panjang")
	} else {
		fmt.Println("Pendek")
	}

}
