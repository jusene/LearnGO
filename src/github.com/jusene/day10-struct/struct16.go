package main

import "fmt"

type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address // 匿名结构体
}

func main() {
	var user2 User
	user2.Name = "JUSENE"
	user2.Gender = "男"
	user2.Address.Province = "浙江" // 通过匿名结构体，字段名访问
	user2.City = "杭州" // 直接访问匿名结构体的字段名

	fmt.Printf("%#v", user2)
}
