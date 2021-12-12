package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(errorVariable bool) {
	defer endApp()
	if errorVariable {
		panic("Aplikasi ERROR") // catch panic of errorVar value is true and send it to recover function which located at the UPPER CODE LINE FROM ERROR PRONE which is on defer function
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	fmt.Println("-----")
	runApp(true)
	fmt.Println("-----")
}
