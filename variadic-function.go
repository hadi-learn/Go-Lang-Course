package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sum := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
	fmt.Println(sumAll(5, 4, 3, 2, 1))

	// slice as an input for parameters
	slice := []int{1, 2, 3, 4, 5}
	total := sumAll(slice...) // three dots required to indicate the slice as input for variadic-function parameter (args in python)
	fmt.Print(total)
}
