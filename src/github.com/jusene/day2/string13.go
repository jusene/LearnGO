package main

import (
	"fmt"
	"unicode"
)


func main() {
	var str = "hello 世界，hello Go语言"
	var count int64
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {  // 检测是否是汉字
			count ++
		}
	}
	fmt.Println(count)
}
