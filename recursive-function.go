package main

import "fmt"

// calculate factorial using loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// calculate factorial using recursive
func factorialRecursive(value int) int {
	// Base Case
	if value == 1 {
		return 1
	} else { // Recursive Case
		return value * factorialRecursive(value-1)
	}
}
func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println("-----")
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
