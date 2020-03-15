package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "123"
	fmt.Printf("orig当前是%T类型，且操作系统是%d位\n", orig, strconv.IntSize)

	// 字符转换int
	num, _ := strconv.Atoi(orig)
	fmt.Printf("%T\n", num)
	// 字符转换float64
	fl, _ := strconv.ParseFloat(orig, 64)
	fmt.Printf("%T\n", fl)
	// 十进制转换字符
	news := strconv.Itoa(num)
	fmt.Printf("%T\n", news)
	// 64位转换字符串
	// strconv.FormatFloat(f float64, fmt bytes, prce int, bitsize int) string
	new2 := strconv.FormatFloat(fl, 'f', 64, 64)
	fmt.Printf("%T\n", new2)
	// fmt表示格式('b','e','f','g'),prce表示精度，bitSize的值为32表示float32，64表示float64
}
