package main

import "fmt"

func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) int {
		return x + y
	}
	ret := add(10, 20)
	fmt.Println(ret)

	// 自执行函数，匿名函数定义加()直接完成
	ret1 := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println(ret1)
}
