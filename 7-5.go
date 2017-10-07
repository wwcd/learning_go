package main

import "fmt"

func main() {
	var sl []int

	fmt.Printf("len = %d, cap = %d\n", len(sl), cap(sl))

	sl = append(sl, 2, 3, 4)

	fmt.Printf("len = %d, cap = %d\n", len(sl), cap(sl))

	s2 := []int{1, 2, 3}

	s3 := s2[1:]

	fmt.Println(s2, s3)
	fmt.Printf("%p, %p\n", &s2[1], &s3)

	s4 := "hello"
	s5 := s4[2:3]
	fmt.Printf("%p, %p\n", &s4, &s5)

	s6 := []byte{1, 2, 3, 4, 5}
	s7 := s6[2:3]
	s7[0] = 32
	fmt.Printf("%p, %p\n", &s6, &s7)
	fmt.Println(s6, s7)
	s7[0] = '1'
	fmt.Println(s6, s7)
}
