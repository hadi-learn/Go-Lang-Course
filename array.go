package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Syukri"
	names[1] = "Hadi"
	names[2] = "Kamil"

	fmt.Println("nama: ", names[0])
	fmt.Println("nama1: ", names[1], "'")
	fmt.Println(names[2])
	// names[3] will just return empty
	fmt.Println("nama3: ", names[3], "'")
	// while names[5] will cause error because index 5 is out of bounds for 5-elements array
	// fmt.Println(names[5])

	// array declaration with values

	var values = [10]int16{
		89, 91, 84,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	// empty index will returning 0
	fmt.Println(values[3])

	// function len(array) will return the length of the array (not the number of data in the array)
	fmt.Println(len(names))
	fmt.Println(len(values))
}
