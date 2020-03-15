package main

import "fmt"

func main() {
	var grade string
	marks := 88

	switch  {
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

	switch  {
	case grade == "A":
		fmt.Println("成绩优秀！")
	case grade == "B":
		fmt.Println("表现良好！")
	case grade == "C", grade == "D" :
		fmt.Println("再接再厉！")
	default:
		fmt.Println("成绩不合格！")
	}

	fmt.Printf("你的成绩为%s\n", grade)
}
