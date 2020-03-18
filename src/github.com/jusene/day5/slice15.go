package main

import "fmt"

func main() {
	slice := make([]int, 1e6)
	fmt.Println(foo(&slice))
}


func foo(slice *[]int) *[]int {
	return slice
}