package main

import "fmt"

type myInt int64 // 自定义类型
type yourInt = int64 // 类型别名

func main() {
	//var a int64
	var n myInt
	n = 100
	//a = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	// fmt.Println(n == a) // 自定义类型是一种全新的类型，类型不同不能比较

	var b int64
	var m yourInt
	m = 100
	b = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println(m == b)
}
