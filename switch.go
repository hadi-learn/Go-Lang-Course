package main

import "fmt"

func main() {
	name := "Syukri"

	switch name {
	case "Syukri":
		fmt.Println("Hi", name)
	case "Hadi":
		fmt.Println("Hello", name)
	default:
		fmt.Println("What's your name?")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Panjang")
	case false:
		fmt.Println("Pendek")
	}

	//switch without condition
	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Kepanjangan")
	case length2 > 5:
		fmt.Println("Panjang")
	default:
		fmt.Println("Pendek")
	}
}
