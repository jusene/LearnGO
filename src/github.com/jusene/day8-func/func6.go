package main

import "fmt"

func main() {
	testLocalVars(20, 2)
}

func testLocalVars(x, y int) {
	fmt.Println(x, y)
	if x > 10 {
		z := 100 // 变量z只在if块中生效
		fmt.Println(z)
	}
	// fmt.Println(z) 无法使用变量z
}