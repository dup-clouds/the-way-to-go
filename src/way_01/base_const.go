package main

import "fmt"

func main() {
	const PI = 3.141592
	println(PI)
	const (
		a, b, c = 1, 2, 3
	)
	fmt.Printf("a=%v,b=%v,c=%v\n", a, b, c)
	const (
		i = iota
		j
		m
	)
	fmt.Printf("i=%v,j=%v,m=%v\n", i, j, m)
	const (
		Apple, Banana     = iota + 1, iota + 2 // Apple=1 Banana=2
		Cherimoya, Durian                      // Cherimoya=2 Durian=3
		Elderberry, Fig                        // Elderberry=3, Fig=4
	)
	fmt.Printf("Apple=%v,Banana=%v,Cherimoya=%v,Durian=%v,Elderberry=%v,Fig=%v\n",
		Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
	const (
		_  = iota             // 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota) // 1 << (10*1) 2**10=1024
		MB                    // 1 << (10*2) 2**20=1048576
		GB                    // 1 << (10*3)
		TB                    // 1 << (10*4)
		PB                    // 1 << (10*5)
		EB                    // 1 << (10*6)
		ZB                    // 1 << (10*7)
		YB                    // 1 << (10*8)
	)
	fmt.Printf("KB=%v,MB=%v,GB=%v,TB=%v,PB=%v,EB=%v\n",
		KB, MB, GB, TB, PB, EB)
}
