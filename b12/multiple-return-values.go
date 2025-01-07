package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c, _ := vals()
	fmt.Println("c=", c)

	_, d := vals()
	fmt.Println("d=", d)

}
