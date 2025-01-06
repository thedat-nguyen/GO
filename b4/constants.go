package main

import (
	"fmt"
	"math"
)

const s string = "constant"
const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println(s)

	const n = 5000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	//
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
