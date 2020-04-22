package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%s会叫\n", d.name)
}

func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	a := dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()

	b := &dog{name: "来福"}
	x = b
	y = b
	x.say()
	y.move()

	c := dog{name: "葫芦"}
	Sayer(c).say()
	Mover(c).move()

	d := &dog{name: "琅琊"}
	Sayer(d).say()
	Mover(d).move()
}
