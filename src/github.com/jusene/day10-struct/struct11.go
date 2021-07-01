package main

import "fmt"

type TwoInt struct {
	a int
	b int
}

func (t *TwoInt) AddThem() int {
	return t.a + t.b
}

func (t *TwoInt) AddToParam(param int) int {
	return t.a + t.b + param
}

func (_ *TwoInt) Echo() {
	fmt.Println("hello world")
}

func main() {
	two := new(TwoInt)
	two.a = 10
	two.b = 20

	fmt.Println(two.AddThem())
	fmt.Println(two.AddToParam(30))
	two.Echo()
}
