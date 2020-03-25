package main

import "fmt"

func main() {
	/*
	var a *int
	*a = 100 // 空指针 panic
	fmt.Println(*a)
	*/
	a := new(int)
	*a = 100
	fmt.Println(*a)

	/*
	var b map[string]int // 空指针 panic
	b["a"] = 100
	fmt.Println(b)
	 */
	b := make(map[string]int)
	b["a"] = 100
	fmt.Println(b)
}
