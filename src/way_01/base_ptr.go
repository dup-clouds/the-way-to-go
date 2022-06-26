package main

import "fmt"

func main() {
	b := 10
	a := &b
	fmt.Println("start=", b)
	fmt.Println("ptr=", &a)
	fmt.Println("a-val=", *a)
	*a = 22
	fmt.Println("end b val=", b)
}
