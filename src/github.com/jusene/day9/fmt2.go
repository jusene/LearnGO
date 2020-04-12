package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "标准输出")
	fmt.Fprint(os.Stderr, "标准错误输出")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错， err:", err)
		return
	}

	name := "jusene"
	fmt.Fprintf(fileObj, "往文件写入信息: %s", name)
}
