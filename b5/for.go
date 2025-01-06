package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for a := range 6 {
		fmt.Println("range", a)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("so le la", n)
	}

}
