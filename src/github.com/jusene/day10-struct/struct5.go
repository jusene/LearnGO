package main

import "fmt"

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	n := test{1, 2, 3, 4}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}
