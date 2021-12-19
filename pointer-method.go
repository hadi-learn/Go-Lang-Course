package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() { // without pointer, data passed to this parameter is pass by value, so it won't affect the origin
	man.Name = "Mr. " + man.Name
	fmt.Println("Value in method:", man.Name)
}

func (man *Man) Nikah() {
	man.Name = "Pak " + man.Name
	fmt.Println("Inside method:", man.Name)
}

func main() {
	syukri := Man{"Syukri"}

	// call struct function without pointer
	syukri.Married()
	fmt.Println(syukri.Name)

	// call struct function with pointer
	syukri.Nikah()
	fmt.Println(syukri.Name)
}
