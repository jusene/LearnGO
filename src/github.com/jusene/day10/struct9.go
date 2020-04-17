package main

import "fmt"

type people struct {
	name, city string
	age int
}

func newPerson(name, city string, age int) *people {
	return &people{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	p := newPerson("jusene", "hangzhou", 27)
	fmt.Printf("%#v", p)
}
