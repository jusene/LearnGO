package main

import "fmt"

func main() {
	fmt.Println("return:", a())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i // 返回值未先声明，匿名需要先声明，再赋值，defer的执行时机在赋值和RET指令间，所以为0
}
