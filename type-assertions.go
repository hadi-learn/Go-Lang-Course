package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	fmt.Println(result)
	fmt.Printf("%T\n", result)
	// fmt.Println(resultString)
	// fmt.Printf("%T\n", resultString)

	// using switch to safely convert type assertions
	switch value := result.(type) {
	case string:
		fmt.Print("value ", value, " is ")
		fmt.Printf("%T\n", value)
	case int:
		fmt.Print("value ", value, " is ")
		fmt.Printf("%T\n", value)
	case bool:
		fmt.Print("value ", value, " is ")
		fmt.Printf("%T\n", value)
	default:
		fmt.Println("value", value, "is unknown type")
	}
}
