package main

import "fmt"

// struct is similar to class in python
type Customer struct { //  struct naming using titlecase for easy reading
	firstName, lastName, address string
	age                          int
}

// normal function with struct data
func sayHi(customer Customer, name string) {
	fmt.Println("Hi", name, "My name is", customer.lastName)
}

// struct method / struct function
func (customer Customer) sayHiStruct(name string) {
	fmt.Println("Hi", name, "My name is", customer.lastName)
}

func (goodbye Customer) sayGoodbye() {
	fmt.Println("Thank you and goodbye", goodbye.firstName, goodbye.lastName)
}

func main() {
	var syukri Customer
	syukri.firstName = "Syukri"
	syukri.lastName = "Kamil"
	syukri.address = "Yogyakarta"
	syukri.age = 35
	fmt.Println(syukri)
	fmt.Println(syukri.firstName)
	fmt.Println(syukri.lastName)
	fmt.Println(syukri.address)
	fmt.Println(syukri.age)

	// other ways to declare a struct
	mutia := Customer{
		firstName: "Mutia",
		lastName:  "Wardani",
		address:   "Yogyakarta",
		age:       32,
	}
	fmt.Println(mutia)
	fmt.Println(mutia.firstName)
	fmt.Println(mutia.lastName)
	fmt.Println(mutia.address)
	fmt.Println(mutia.age)

	hadi := Customer{"Hadi", "Kamil", "Yogyakarta", 35} // not recommended because if there is a change in Customer struct type, this will prone error
	fmt.Println(hadi)

	// calling method/function the normal way
	sayHi(syukri, "Mutia")
	// calling a struct method/function
	mutia.sayHiStruct("Syukri")
	mutia.sayGoodbye()
}
