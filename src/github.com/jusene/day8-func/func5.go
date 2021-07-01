package main

import "fmt"

// 定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Println(num)
}

func main() {
	testNum() // 当局部变量和全局变量重名，优先使用局部变量
}
