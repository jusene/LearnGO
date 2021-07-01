package main

import "fmt"

func main() {
	nums := []int{}
	renums := make([]int, 0)
	var anums []int
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("renums: %v, len: %d, cap: %d\n", renums, len(renums), cap(renums))
	fmt.Printf("anums: %v, len: %d, cap: %d\n", anums, len(anums), cap(anums))

	if nums == nil {
		fmt.Println("nums is nil")
	}
	if renums == nil {
		fmt.Println("renums is nil")
	}
	if anums == nil {
		fmt.Println("anums is nil")
	}
}
