package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Syukri",
		"city": "Jogja",
	}
	fmt.Println(person)
	fmt.Println("name:", person["name"])
	fmt.Println("city of origin:", person["city"])

	person["title"] = "Software Engineer"
	fmt.Println(person)

	person["title"] = "Software Developer"
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["author"] = "Syukri"
	book["title"] = "Learn Go-Lang"
	book["year"] = "2021"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
