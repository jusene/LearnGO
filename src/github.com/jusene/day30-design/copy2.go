package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))

	dnums := nums[0:2]
	dnums = append(dnums, []int{2, 3}...)
	dnums[0] = 100
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("nums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
}
