package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	name string
}

func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

func (c car) move() {
	fmt.Printf("%s会跑\n", c.name)
}

func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{name: "路虎"}
	x = a
	x.move()
	x = b
	x.move()
}
