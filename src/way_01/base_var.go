package main

import (
	"fmt"
	"os"
	"runtime"
)

var a = "G"
var b string

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println()
	n()
	m()
	n()
	fmt.Println()
	b = "G"
	print(b)
	f1()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}

func f1() {
	b := "O"
	print(b)
	f2()
}

func f2() {
	print(b)
}
