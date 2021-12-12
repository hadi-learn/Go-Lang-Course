package main

import "fmt"

func main() {
	counter := 0
	name := "Syukri"

	increment := func() {
		name := "Hadi" // this variable name will not accessible outside the scope of the inner function
		fmt.Println(name)
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
