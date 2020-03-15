package main

import "fmt"

func main() {
	var grade string
	switch marks := 90; { // 分号不能省略
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("你的成绩为%s\n", grade)
}
