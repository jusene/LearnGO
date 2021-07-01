package main

import "fmt"

func main() {
	var a []string   // 声明一个字符串切片
	var b = []string{}  // 声明一个整型切片并初始化
	slice := make([]string, 0)
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(slice == nil) // false
}
