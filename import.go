package main

import (
	"Go-Lang-Course/helper"
	"fmt"
)

func main() {
	helper.SayHello("Syukri")
	// helper.sayGoodbye("Syukri") --- tidak bisa di import karena dimulai dengan huruf kecil
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) --- tidak bisa di akses karena dimulai dengan huruf kecil
}
