package main

import (
	"fmt"
)

func main() {
	age := ageMinOrMax("min", 1, 3, 4, 59)
	fmt.Printf("年龄最小的参数是%d岁\n", age)

	ageArr := []int{6, 4, 43, 2, 32}
	age = ageMinOrMax("max", ageArr...)
	fmt.Printf("年龄最大的参数是%d岁\n", age)
}

func ageMinOrMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}

	if m == "max" {
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	} else if m == "min" {
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	} else {
		e := -1
		return e
	}

}
