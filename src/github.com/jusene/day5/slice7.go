package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	/* 切片不能比较
	if a == b {
		fmt.Println('true')
	}
	 */
	fmt.Println(a, b)

	c := [5]int{1, 2, 3, 4, 5}
	d := [5]int{1, 2, 3, 4, 5}
	// 数组可以直接比较
	if c == d {
		fmt.Println("true")
	}
}
