package main

import "fmt"

var i = 0

func print(i int) {
	fmt.Println(i, &i)
}

func main() {
	for ; i < 5; i++ {
		defer print(i)
	}
}
