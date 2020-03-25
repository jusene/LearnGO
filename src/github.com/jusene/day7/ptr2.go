package main

import "fmt"

func main() {
	a := 10
	modify1(10)
	fmt.Println(a)
	modify2(&a) // 指针操作可以改变原来的值
	fmt.Println(a)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}