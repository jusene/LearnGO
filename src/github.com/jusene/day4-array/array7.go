package main

import "fmt"

func main() {
	nums := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
