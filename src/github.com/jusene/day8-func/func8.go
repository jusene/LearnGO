package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	ret := calc(10, 20, add)
	fmt.Println(ret)
}
