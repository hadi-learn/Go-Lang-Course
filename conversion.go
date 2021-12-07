package main

import "fmt"

func main() {
	var nilai32 int32 = 512
	var nilai64 int64 = int64(nilai32)
	var nilai16 = int16(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	fmt.Printf("Var nilai32 = %T\n", nilai32)
	fmt.Printf("Var nilai16 = %T\n", nilai16)
	fmt.Printf("Var nilai64 = %T\n", nilai64)
	fmt.Printf("Var nilai8 = %T\n", nilai8)

	var name = "Syukri"
	var e byte = name[0]
	var eString string = string(e)
	var char = string(name[0])
	firstChar := string(name[0])

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	fmt.Println(char)
	fmt.Println(firstChar)
}
