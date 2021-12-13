package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi nol tidak diizinkan")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError error = errors.New("Ups error")
	fmt.Println(contohError.Error())
	fmt.Println(Pembagian(10, 0))
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
}
