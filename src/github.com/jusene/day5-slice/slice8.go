package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
