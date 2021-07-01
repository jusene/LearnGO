package main

import "fmt"

func Add() func(b int) int {
	var a int
	return func(b int) int {
		a += b
		return a
	}
}

func main() {
	var f = Add()
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
}
