package main

import "fmt"

func main() {
	// 十进制
	var a int64 = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 二进制 1010
	fmt.Printf("%o \n", a) // 八进制 12
	fmt.Printf("%x \n", a) // 十六进制 a

	// 八进制 以0开头
	var b int64 = 077
	fmt.Printf("%d \n", b) // 十进制 63
	fmt.Printf("%b \n", b) // 二进制 111111
	fmt.Printf("%o \n", b) // 77
	fmt.Printf("%x \n", b) // 十六进制 3f

	// 十六进制 以0x开始
	var c int64 = 0xff
	fmt.Printf("%d \n", c) // 十进制 255
	fmt.Printf("%b \n", c) // 二进制 11111111
	fmt.Printf("%o \n", c) // 八进制 377
	fmt.Printf("%x \n", c) // ff
}
