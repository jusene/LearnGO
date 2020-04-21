package main

import "fmt"

type Sayer interface {
	say()
	move()
}

type cat struct {}

func (c cat) say() {
	fmt.Println("喵喵")
}

func (c cat) move() {
	fmt.Println("猫动了")
}

type dog struct {}

func (d dog) say() {
	fmt.Println("汪汪")
}

func (d dog) move() {
	fmt.Println("狗动了")
}

func main() {
	c := cat{}
	d := dog{}
	Sayer(c).say()
	Sayer(c).move()
	Sayer(d).say()
	Sayer(d).move()
}
