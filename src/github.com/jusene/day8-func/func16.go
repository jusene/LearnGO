package main

import "fmt"

func main() {
	// 阶乘
	var i int = 15
	fmt.Printf("%d的阶乘是%d", i, Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if (n > 0) {
		result = n * Factorial(n - 1)
		return result
	}
	return 1
}