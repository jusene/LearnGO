package main

import "fmt"

func main() {
	fmt.Println("打开文件")
	defer fmt.Println("关闭远程连接")
	defer fmt.Println("关闭文件")
	fmt.Println("读取文件")
}
