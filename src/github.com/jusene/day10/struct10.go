package main

import "fmt"

// Person 结构体
type Person struct {
	name string
	age int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person 的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

func main() {
	p1 := NewPerson("jusene", 27)
	p1.Dream()
}
