package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Eko")
	fmt.Println(person)
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		for key, value := range person {
			fmt.Println(key, ":", value)
		}
	}
	person2 := NewMap("")
	fmt.Println(person2)
	if person2 == nil {
		fmt.Println("Data kosong")
	} else {
		for key, value := range person2 {
			fmt.Println(key, ":", value)
		}
	}
}
