package main

import "fmt"

func Add(a int) func(b int) int {
	return func(b int) int {
		a += b
		return a
	}
}

func main() {
	var f = Add(10)
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
}
