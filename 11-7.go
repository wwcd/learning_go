package main

import "fmt"
import "./sort"

func ints() {
	data := []int{74, 59, -784, 7586}
	a := sort.IntArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}

	fmt.Printf("The sorted array is: %v\n", a)
}

func strings() {
	data := []string{"monday", "friday", "tuesday"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}

	fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{1, "TUE", "Tuesday"}

	data := []*day{&Tuesday, &Monday, &Sunday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}

	fmt.Printf("The sorted array is: %v\n", a)
}

func main() {
	ints()
	strings()
	days()
}
