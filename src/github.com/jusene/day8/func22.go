package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() { // 有效，在defer语句的匿名函数中调用
		fmt.Println(recover())
	}()

	defer func() { // 无效，间接调用recover(),返回nil
		func() {
			fmt.Println(recover())
		}()
	}()

	defer fmt.Println(recover()) /* 无效，recover() 相当于直接调用然后被外部函数打印，返回nil*/
	defer recover() // 无效，相当于直接调用recover()，返回nil
	panic("发生错误")
}
