package main

import "fmt"

func main() {
	fmt.Println("return:", b())
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()

	return i // 返回值有名，直接赋值，匿名函数的最大特点可以继承变量的值，RET返回之前需要执行defer，所以i=2
}
