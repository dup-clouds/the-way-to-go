package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "a|b|c"
	str2 := "aa bb cc"
	slice1 := strings.Split(str1, "|")
	slice2 := strings.Fields(str2)
	fmt.Println("\nprint slice1 start")
	for _, v := range slice1 {
		fmt.Print(v + " ")
	}
	fmt.Println("\nprint slice1 end")
	fmt.Println("\nprint slice2 start")
	for _, v := range slice2 {
		fmt.Print(v + " ")
	}
	fmt.Println("\nprint slice2 end")
	fmt.Println(strings.Join(slice1, ";"))
	fmt.Println(strings.Join(slice2, ";"))
	newStr := strings.Repeat("clouds", 3)
	fmt.Println(newStr)
}
