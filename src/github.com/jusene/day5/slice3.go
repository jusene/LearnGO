package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice1 := slice[2:5]
	slice2 := slice1[1:3]
	slice2[0] = 100
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
