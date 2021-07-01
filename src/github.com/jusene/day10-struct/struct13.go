package main

import "fmt"

type People struct {
	name string
	age int
}

// 指针类型接受者
func(p *People) SetAge(newAge int) {
	p.age = newAge
}

// 值类型的接受者
func(p People) SetAge2(newAge int) int {
	p.age = newAge
	return p.age
}


func main() {
	per := People{"jusene", 27}
	fmt.Println(per.age) // 27
	per.SetAge(30)
	fmt.Println(per.age) // 30

	newAge := per.SetAge2(40)
	fmt.Println(per.age) // 30
	fmt.Println(newAge) // 40
}
