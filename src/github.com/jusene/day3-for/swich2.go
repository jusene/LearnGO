package main

import "fmt"

func main() {
	var grade string
	switch marks := 90; marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 60:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("你的成绩为%s\n", grade)
}

