package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	newNums := append(nums[:2], nums[3:]...)
	fmt.Println(newNums)
}
