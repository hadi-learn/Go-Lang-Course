package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello there"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("")
	fmt.Println(result)
	fmt.Println(getHello("Syukri"))
}
