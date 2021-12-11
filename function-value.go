package main

import "fmt"

// function as a value (maybe similar to class in python)
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// function as a parameter to other function (maybe similar to decorator in python)
func sayHelloWithFilter(name string, filter func(string) string) {
	nameToFilter := name
	fmt.Println("Hello ", filter(nameToFilter))
}

// using alias to function to simplify parameters
type Filter func(string) string

func sayHelloWithFilterAlias(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "bad" {
		return "..."
	} else {
		return name
	}
}

func main() {
	result := getGoodBye // function as a value doesn't need parentheses ()
	fmt.Println(result("Syukri"))

	fmt.Println("-----")

	sayHelloWithFilter("Syukri", spamFilter)
	sayHelloWithFilter("bad", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Hadi", filter)

	fmt.Println("-----")

	sayHelloWithFilterAlias("Syukri", spamFilter)
	sayHelloWithFilterAlias("bad", spamFilter)
	sayHelloWithFilterAlias("Hadi", filter)
	sayHelloWithFilterAlias("bad", func(s string) string {
		if s == "bad" {
			return "..."
		} else {
			return s
		}
	})
}
