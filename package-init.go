package main

import (
	// _ "Go-Lang-Course/database" // character "_" is used as blank identifier so the code will ignore unused import
	"Go-Lang-Course/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
