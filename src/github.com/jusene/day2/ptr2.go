package main

import "fmt"

func main() {
	var null *int
	nu := &null
	fmt.Printf("null --> nil: %x\n", null)
	fmt.Printf("null的内存地址: %x\n", nu)
	fmt.Printf("nu --> null --> nil: %x\n", *nu)
	fmt.Printf("nu的内存地址：%x\n", &nu)


	A := 10
	AP := &A
	APP := &AP
	fmt.Printf("A: %d\n", A)
	fmt.Printf("AP: %x\n", AP)
	fmt.Printf("*AP: %d\n", *AP)
	fmt.Printf("APP: %x\n", APP)
	fmt.Printf("*APP: %x\n", *APP)
	fmt.Printf("**APP: %d\n", **APP)
}