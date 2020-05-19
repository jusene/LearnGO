package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test")
	if err != nil {
		panic(err)
	}
	file.WriteString("hello golang")
	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5

	// 用来计算offset的初始位置
	// 0 文件开始位置
	// 1 当前位置
	// 2 文件结尾数
	whence := 0

	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		panic(err)
	}

	fmt.Println("移到位置5: ", newPosition)

	// 从当前位置回退两字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("回退两个字节: ", newPosition)

	// 获取当前位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("当前: ", currentPosition)

	// 移到文件开头
	newPosition, err = file.Seek(0, 0)
	fmt.Println("开头", newPosition)
}
