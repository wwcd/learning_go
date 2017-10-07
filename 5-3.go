package main

import "fmt"

func main() {
	num1 := 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
	case 101:
		fmt.Println("It's equal to 98 or 100")
	}
}
