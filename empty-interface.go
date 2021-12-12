package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)

	result := Ups(2)
	fmt.Println(result)

	get_data := Ups(3)
	fmt.Println(get_data)
}
