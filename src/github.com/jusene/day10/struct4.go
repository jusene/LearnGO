package main

import "fmt"

type person struct {
	name string
	age int
	city string
}

func s1() person {
	p := person{
		name: "jusene",
		age:  27,
		city: "hangzhou",
	}
	return p
}

func s2() *person {
	p := &person{
		name: "jusene",
		age:  27,
		city: "hangzhou",
	}
	return p
}

func s3() *person {
	p := &person{
		"jusene",
		27,
		"hangzhou",
	}
	return p
}

func main() {
	p := s1()
	p1 := s2()
	p2 := s3()
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)
}
