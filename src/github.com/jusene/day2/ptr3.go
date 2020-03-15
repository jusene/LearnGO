package main

import "fmt"

const MAX int = 3

func main() {

	a := []int{10,100,200}
	var ptr [MAX]*int

	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("a[%d]的地址 = %x\n", i, ptr[i] )
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d]的值是: %d\n", i, *ptr[i])
	}
}