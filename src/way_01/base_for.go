package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		var v int
		// 0 0 0 0 0
		fmt.Printf("%d ", v)
		v = 5
	}
	// 无限打印
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	// 无限打印
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

}
