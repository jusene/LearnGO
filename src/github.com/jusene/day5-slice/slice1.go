package main

import "fmt"

func main() {
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
}
