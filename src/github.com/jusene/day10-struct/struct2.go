package main

import "fmt"

type person struct {
	name string
	age int
	city string
}

func s1() person {
	var p person
	p.name = "jusene"
	p.age = 27
	p.city = "hangzhou"

	fmt.Printf("%#v\n", p)
	return p
}

func s2() *person {
	s := new(person)
	s.name = "jusene" // 底层 (*s).name = "jusene" Go语言帮我们实现的语法糖
	s.age = 27
	s.city = "hangzhou"

	fmt.Printf("%#v\n", s)
	return s
}

func s3() *person {
	t := &person{
		name: "jusene",
		age:  27,
		city: "hangzhou",
	}

	fmt.Printf("%#v\n", t)
	return t
}


func main() {
	s1()
	s2()
	s3()
}
