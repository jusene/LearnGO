package main

import "fmt"

func main() {
	str := "jusene"
	for i, char := range str {
		fmt.Printf("字符串第%d个字符串的值为%d\n",i, char)
	}
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
	numbers := []int{1, 2, 3}
	for i, x := range numbers {
		fmt.Println(i, x)
	}
	nums := [5]int{1,2,3,4}
	for i, x := range nums {
		fmt.Println(i, x)
	}
}
