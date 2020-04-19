package main

import "fmt"

type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

func main() {
	user1 := User{
		Name:    "jusene",
		Gender:  "男",
		Address: Address{
			Province: "浙江",
			City: "杭州",
		},
	}

	fmt.Printf("%#v", user1)
}
