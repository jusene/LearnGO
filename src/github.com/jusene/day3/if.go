package main

import "fmt"

func main() {
	a := 13
	if a > 20 {
		fmt.Printf("a大于20\n")
	} else if a < 10 {
		fmt.Printf("a小于10\n")
	} else if a == 11 {
		fmt.Printf("a等于11\n")
	} else {
		fmt.Printf("a大于10\n")
		fmt.Printf("a小于20\n")
		fmt.Printf("a不等于11\n")
	}
}


