package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Yogyakarta", "D.I. Yogyakarta", "Indonesia"}
	address2 := address1 // pass by value or deep copy

	address2.City = "Sleman"

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 Address = Address{"Bantul", "D.I. Yogyakarta", "Indonesia"}
	var address4 *Address = &address3 // pass by reference or shallow copy indicated by "&" symbol
	fmt.Println(address3)

	address4.City = "Wonosari"
	fmt.Println(address3)
	fmt.Println(address4)

	// when creating a variable using pointer to a struct instead of to other variable using same struct
	// Go will create another new data on memory so
	// any changes on that variable won't affects other variable with the same struct

	address4 = &Address{"Jakarta", "D.K.I. Jakarta", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address4)

	// to be able to switch the data pointer of all variables that using the same struct to a new data created by other variable, Go provide Operator "*"

	address5 := Address{"Yogyakarta", "D.I. Yogyakarta", "Indonesia"}
	address6 := &address5
	address7 := &address5
	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)
	address6.City = "Bantul"
	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)
	*address6 = Address{"Jakarta", "D.K.I. Jakarta", "Indonesia"}
	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)

	// Function new can be used to create an empty struct variable
	var address8 *Address = new(Address)
	fmt.Println(address8)
	address8.City = "Yogyakarta"
	fmt.Println(address8)
}
