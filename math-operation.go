package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	a := 10
	b := 10
	c := a * b
	fmt.Println(c)

	d := a / 5
	fmt.Println(d)

	// this will result the floor division instead of floating number because the data types are integer
	// for floor division, only float64 data type can be done with the math.Floor(float64(a)/float64(b)) function
	e := a / 3
	fmt.Println(e)
	fmt.Printf("e is: %T\n", e)

	f := float32(a) / 3
	fmt.Println(f)
	fmt.Printf("f is: %T\n", f)

	g := 10 % 3
	fmt.Println(g)

	var i int8 = 10
	i += 3
	fmt.Println(i)
	i *= 2
	fmt.Println(i)

	// unary operator

	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

}
