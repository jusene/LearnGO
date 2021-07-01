package main

import "fmt"

func main() {

	// 遍历决定处理第几行
	for y:=1; y <= 9; y++ {

		// 遍历。决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%dx%d=%d ", x, y, x*y)
		}

		fmt.Println()
	}
}
