package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["张三"] = 90
	dictMap["小明"] = 100

	v, ok := dictMap["张三"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("没这个人")
	}
}
