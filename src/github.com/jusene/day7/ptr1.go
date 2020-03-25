package main

import "fmt"

func main() {
	// 指针取值
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
