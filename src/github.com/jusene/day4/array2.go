package main

import "fmt"

func main() {
	array := [5]*int{1: new(int)}
	*array[1] = 10
	fmt.Println(array)
	fmt.Println(*array[1])
}
