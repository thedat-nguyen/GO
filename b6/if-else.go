package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 la so chan")
	} else {
		fmt.Println("7 la so le")
	}

	if 7%4 == 0 {
		fmt.Println(" 8 chia het cho 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("1 trong 2dk dung se in ra")
	}
	//
	if num := 19; num < 0 {
		fmt.Println(num, "so am")
	} else if num < 10 {
		fmt.Println(num, "co 1 ki tu so")
	} else {
		fmt.Println(num, "co nhieu ki tu so")
	}
	//
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
	//
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}

// Ít hơn<
// Nhỏ hơn hoặc bằng<=
// Lớn hơn>
// Lớn hơn hoặc bằng>=
// Bằng với==
// Không bằng!=

// Logic VÀ&&
// Logic HOẶC||
// Logic KHÔNG!
