package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop number:", i)
	}
	// using break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("loop number:", i)
	}
	// using continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("loop number:", i)
	}
}
