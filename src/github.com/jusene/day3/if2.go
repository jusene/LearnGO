package main

import "fmt"

func main() {
	if a := 10; a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a的值是: %d\n", a)
	}
	//fmt.Println(a)
}
