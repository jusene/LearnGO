package main

import "fmt"

type author struct {
	name string
	age int
	tag []string
}

func main() {
	var s author
	// s := author

	s.name = "jusene"
	s.age = 27
	s.tag = []string{"devops", "ops"}

	fmt.Println(s)
}
