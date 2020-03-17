package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newslice := append(slice, 100, 200)
	fmt.Println(newslice)


	slice1 := []int{10, 20, 30}
	newslice1 := append(slice, slice1...)
	fmt.Println(newslice1)
}
