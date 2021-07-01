package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 6)
	copy(c, a)
	fmt.Println(a)
	fmt.Println(c)
	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)
}
