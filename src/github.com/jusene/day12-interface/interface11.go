package main

import "fmt"

func show(a interface{}) (b interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
	return true
}
func main() {
	a := 21212
	x := show(a)
	fmt.Printf("type:%T value:%v\n", x, x)
	b := false
	show(b)
}
