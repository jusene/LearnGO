package main

import "fmt"

func main() {
	a := 0
	b := 5
	for a < b {
		a++
		fmt.Printf("a的值是：%d\n", a)
	}
}
