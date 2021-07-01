package main

import "fmt"

func main() {
	a := []int{}
	fmt.Println(len(a), a == nil)

	var b []int
	fmt.Println(len(b), b == nil)

	c := make([]int, 0, 5)
	fmt.Println(len(c), c == nil)

}
