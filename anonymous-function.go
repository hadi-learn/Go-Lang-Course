package main

import "fmt"

type Blacklist func(string) bool

// type Filter func(string) string

// func filterName(name string) string {
// 	if name == "bad" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your access denied", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	checkname := func(name string) bool {
		return name == "bad"
	}

	registerUser("Syukri", checkname)
	registerUser("bad", checkname)
	registerUser("Hadi", func(s string) bool {
		return s == "bad"
	})
	registerUser("bad", func(s string) bool {
		return s == "bad"
	})

}
