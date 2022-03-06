package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func SayHello(name string) {
	fmt.Println("Hello,", name)
	// return "Hello, " + name
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye,", name)
}
