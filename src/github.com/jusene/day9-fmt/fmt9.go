package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Print("请输入用户:")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Print("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Print("请输入婚姻:")
	fmt.Scanf("%t", &married)
	fmt.Println()

	fmt.Printf("name:%s age:%d married:%t \n", name, age, married)
}
