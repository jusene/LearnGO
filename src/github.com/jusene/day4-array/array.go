package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println(array)
	array1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array1)
	// 编译器根据初始值的个数自行推断数组的长度
	array2 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("%T\n", array2)
	fmt.Println(array2)
	array3 := [5]int{1: 10, 3: 20}
	fmt.Println(array3)
}
