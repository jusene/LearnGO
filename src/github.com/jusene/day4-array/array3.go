package main

import "fmt"

func main() {
	var array [4][2]int
	fmt.Println(array)

	// 声明并初始化外层数组中索引为1和3的元素
	array1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array1)

	// 声明并初始化外层数组和内层数组的单个元素
	array2 := [4][2]int{1: {0: 21}, 3: {1: 41}}
	fmt.Println(array2)
}
