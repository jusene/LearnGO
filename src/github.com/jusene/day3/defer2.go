package main

import "fmt"

var i = 0

func print() {
	fmt.Println(i, &i)
}

func main() {
	for ; i < 5; i++ {
		defer print()
	}
}
