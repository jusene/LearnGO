package main

import "fmt"

func main() {
	x := 3
	y := 4
	a, b := SumAndProduct(x, y)
	fmt.Println(a, b)
	c, d := SumAndProduct1(x, y)
	fmt.Println(c, d)
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

// 命名的返回值就相当于在函数的声明的时候声明了一个变量
func SumAndProduct1(A, B int) (add int, mul int) {
	add = A + B
	mul = A * B
	return
}