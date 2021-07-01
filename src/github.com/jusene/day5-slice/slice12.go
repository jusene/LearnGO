package main

import "fmt"

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		a = append(a, fmt.Sprintf("%v", i))
		fmt.Println(a)
	}
	fmt.Println(a)
}
