package main

import "fmt"

func main() {
	// 数组
	nums := [5]int{}
	nums[0] = 1
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))

	// 切片
	dnums := nums[1:2]
	dnums[0] = 3
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
}
