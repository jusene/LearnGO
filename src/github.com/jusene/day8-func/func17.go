package main

import "fmt"

func main() {
	// 斐波那契数列
	for i := 0; i < 10; i++  {
		fmt.Printf("%d\t", Fibonacci(i))
	}
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
