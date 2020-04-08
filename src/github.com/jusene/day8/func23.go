package main

import "log"

func main() {
	test()
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("捕获到的异常: %v", r) // recover()只能捕获最近的一个异常
		}
	}()

	defer func() {
		panic("第二个错误")
	}()

	panic("第一个错误")
}
