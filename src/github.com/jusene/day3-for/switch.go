package main

import "fmt"

func main() {
	var grade string
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Printf("你的成绩为%s\n", grade)
}
