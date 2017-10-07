package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func f(a [5]int) {
	fmt.Println(a)
}

func fp(a *[5]int) {
	fmt.Println(a)
}

func main() {
	var arr1 [5]int
	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	f(arr1)
	fp(&arr1)
}
