package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func main() {
	var c calculation // 申明一个calculation类型的变量
	c = add
	fmt.Printf("%T", c) // main.calculation
	fmt.Println(c(1, 3))

	f := add          // 将add函数赋值给变量f
	fmt.Printf("%T", f) // func(int, int) int
	fmt.Println(f(1, 3))
}
