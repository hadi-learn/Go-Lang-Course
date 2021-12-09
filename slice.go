package main

import "fmt"

func main() {
	var months = [...]string{ // the use of [...] indicate array with undecided size
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
	newSlice := make([]string, 2, 5) // the use of [] indicate slice with undecided size
	newSlice[0] = "Syukri"
	newSlice[1] = "Hadi"
	fmt.Println(newSlice)

	newSlice2 := make([]string, 4, 5)
	// adding data one at a time to newSlice then changing it's data won't affect the array's origin
	newSlice2[0] = months[2]
	newSlice2[1] = months[3]
	newSlice2[2] = months[4]
	fmt.Println(newSlice2)
	newSlice2[0] = "Ganti"
	fmt.Println(newSlice2)
	fmt.Println(months)

	// adding data with slice syntax to newSlice then changing it's data will affect the array's origin
	newSlice3 := make([]string, 4, 5)
	newSlice3 = months[5:8]
	fmt.Println(newSlice3)
	newSlice3[0] = "Ganti"
	fmt.Println(newSlice3)
	fmt.Println(months)

	// copy slice (make sure the size is the same)
	copy_slice := make([]string, len(newSlice), cap(newSlice))
	copy(copy_slice, newSlice)
	fmt.Println(copy_slice)
	fmt.Println(newSlice)

	// different between array and slice
	iniArray := [5]int8{1, 2, 3, 4, 5}
	iniArray2 := [9]int8{1, 2, 3, 4, 5, 6}
	iniArray3 := [...]int8{1, 2, 3, 4, 5, 6, 7, 8}
	iniSlice := []int8{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniArray3)
	fmt.Println(iniSlice)

}
