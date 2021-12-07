// aliasing data type

package main

import "fmt"

func main() {
	type NomorKTP string
	type Married bool

	var ktpSyukri NomorKTP = "1230984756"
	var ktpMutia NomorKTP = "2349870165"

	var marriedSyukri Married = true
	var marriedMutia Married = true

	fmt.Println(ktpSyukri)
	fmt.Println(ktpMutia)
	fmt.Println(marriedSyukri)
	fmt.Println(marriedMutia)
}
