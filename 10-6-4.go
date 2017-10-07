package main

import "fmt"
import "./person"

func main() {
	p := new(person.Person)

	// p.firstName = "Eric"

	p.SetFirstName("Eric")

	fmt.Println(p.FirstName())
}
