package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 2 {
				// 跳到标签
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}

	breakTag:
		fmt.Println("结束两层for循环")
}

