package main

import "os"

func main() {
	file, _ := os.Create("test")
	file.WriteString("hello world")
	file.Close()
	// err := os.Remove("test")
	err := os.Truncate("test", 1) // 传入0会清空文件
	if err != nil {
		panic(err)
	}
}
