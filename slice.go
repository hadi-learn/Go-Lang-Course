package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",   //0
		"Februari",  //1
		"Maret",     //2
		"April",     //3
		"Mei",       //4
		"Juni",      //5
		"Juli",      //6
		"Agustus",   //7
		"September", //8
		"Oktober",   //9
		"November",  //10
		"Desember",  //11
	}
	fmt.Println(months[0])
	fmt.Println(len(months))
	fmt.Println(cap(months))

	// Slice is a shallow-copy so be careful!!
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[6:10]
	fmt.Println(slice2)
	fmt.Println(slice2[0])
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// append slice
	var slice3 = months[9:]
	fmt.Println(slice3)
	fmt.Println(slice3[0])
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	var new_slice3 = append(slice3, "Tambah")
	fmt.Println(slice3)
	fmt.Println(new_slice3)
	new_slice3[0] = "Oktober-false"
	fmt.Println(slice3)
	fmt.Println(new_slice3)

	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println(cap(months))

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Syukri"
	newSlice[1] = "Hadi"
	fmt.Println(newSlice)

	newSlice2 := make([]string, 1, 5)
	newSlice2[0] = months[2]
	//newSlice2[1] = months[3]
	//newSlice2[2] = months[4]
	fmt.Println(newSlice2)
	newSlice2[0] = "Ganti"
	fmt.Println(newSlice2)
	fmt.Println(months)

}
