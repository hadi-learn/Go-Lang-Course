package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	var lulusUjian bool = ujian > 80
	var lulusAbsensi bool = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println("Apakah lulus?: ", lulus)
}
