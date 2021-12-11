package main

import (
	"fmt"
)

func main() {
	// basic loop
	for count := 1; count <= 10; count++ {
		fmt.Println("Count number:", count)
	}

	// while like loop
	counter := 1
	for counter < 10 {
		fmt.Println("Loop number:", counter)
		counter++
	}

	slice := []string{"Syukri", "Hadi", "Kamil"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range (like enumerate in python) in array/slice
	names := []string{"Mutia", "Kusuma", "Wardani"}
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
		fmt.Println("Index:", index, "Value:", value)
	}
	for _, values := range names {
		fmt.Println(values)
	}

	// for range in map
	person := make(map[string]string)
	person["name"] = "Syukri"
	person["title"] = "Software Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
