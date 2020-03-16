package main

import "fmt"

func main() {
	var a = [...]string{"北京", "上海", "广州", "深圳", "杭州"}
	// for循环遍历
	for i:=0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
