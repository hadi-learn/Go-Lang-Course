package main

import "fmt"

func printHello() {
	fmt.Println("Hello")
}

func main() {
	printHello()
	for i := 0; i <= 10; i++ {
		printHello()
	}
}
