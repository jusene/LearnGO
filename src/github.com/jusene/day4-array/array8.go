package main

import "fmt"

func main() {
	nums := [...]int{1, 3, 4, 7, 8}
	for i, _ := range nums {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == 8 {
				fmt.Printf("(%d %d)\n", i, j)
			}
		}
	}
}
