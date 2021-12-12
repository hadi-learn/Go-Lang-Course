package main

import "fmt"

// common interface syntax
type hasName interface {
	getName() string
}

// interface can be used as parameter for a function
func sayHello(hasname hasName) {
	fmt.Println("Hello", hasname.getName())
}

// interface implementation example 1
type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

// interface implementation example 2
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	// example 1 execution
	person := Person{
		Name: "Syukri",
	}
	sayHello(person)

	var member Person
	member.Name = "Mutia"
	sayHello(member)

	// example 2 execution
	cat := Animal{
		Name: "Kucing",
	}
	sayHello(cat)

	var ayam Animal
	ayam.Name = "Chicken"
	sayHello(ayam)
}
