package main

import "fmt"

func main() {
	num := 100
	filter(isBig, num)
}

type funcBool func(x int) bool

func filter(f funcBool, num int) bool {
	if f(num) {
		fmt.Println("ok")
		return true
	} else {
		fmt.Println("no")
		return false
	}
}

func isBig(x int) bool {
	if x >= 100 {
		return true
	}
	return false
}