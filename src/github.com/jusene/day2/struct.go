package main

import (
	"fmt"
)

func main() {
	type stu struct {
		Name string
		Age int64
	}

	student := stu{
		Name: "guoxing",
		Age: 12,
	}

	fmt.Println(student.Name, student.Age)
}