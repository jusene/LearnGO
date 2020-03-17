package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	// slice[i:j:k] 对底层容量k的切片，长度为j-i，容量为k-i
	newSlice := slice[2:3:4]
	fmt.Printf("%v", newSlice)
	fmt.Println(len(newSlice), cap(newSlice))
}
