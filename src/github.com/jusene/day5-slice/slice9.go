package main

import "fmt"

func main() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p \n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
