package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("i = 2 vay nen in ra case ", i, "!!!")
	switch i {
	case 1:
		fmt.Println("case 1 nd la 222")
	case 2:
		fmt.Println("case 2 nd la 123 ")
	case 3:
		fmt.Println("case 3 nd la 336")
	}
	//
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("cuoi tuan")
	default:
		fmt.Println("1 ngay trong tuan")
	}
	//
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("trc buoi trua")
	default:
		fmt.Println("sau buoi trua")
	}
	//
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("true là một giá trị kiểu bool")
		case int:
			fmt.Println("in ra case int ")
		case string:
			fmt.Println("in ra case kieu chu")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
