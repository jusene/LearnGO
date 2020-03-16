package main

import "fmt"

var array [1e6]int

func main() {
	foo(&array)
}

func foo(array *[1e6]int) {
	fmt.Println(array)
}
