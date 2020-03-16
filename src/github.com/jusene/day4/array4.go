package main

import "fmt"

func main() {
	a := [3][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	b := [...][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	// 不支持多维数组的内层使用
	/*
	c := [3][...]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	*/

}
