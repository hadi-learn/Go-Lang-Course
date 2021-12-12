package main

import "fmt"

func logging() {
	fmt.Println("Finished call a function!")
}

// using defer to call a function will ignore error and continue call the function
func runFunction(value int) {
	defer logging()
	fmt.Println("Running a function.")
	result := 10 / value // if value is 0 will resulting error dividing by zero
	fmt.Println(result)
}

// without defer to call a function will catch error
func runApp(value int) {
	fmt.Println("Running an app.")
	result := 10 / value // if value is 0 will resulting error dividing by zero
	fmt.Println(result)
	logging()
}

func main() {
	runFunction(0)
	// runApp(0)
}
