package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{"jusene", 27}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)
}
