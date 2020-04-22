package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵")
}

func (c cat) move() {
	fmt.Println("动了")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
