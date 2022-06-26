package main

import "fmt"

func main() {
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
	fmt.Printf("%03d\n", n)
	fmt.Printf("%5.4g\n", 3.14159)
	type Rope string
	var a Rope = "str"
	fmt.Println(a)
}
