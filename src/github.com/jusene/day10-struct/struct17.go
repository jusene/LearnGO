package main

import (
	"fmt"
	"time"
)

type Address struct {
	Province string
	City string
	CreateTime time.Time
}

type Email struct {
	Account string
	CreateTime time.Time
}

type User struct {
	Name string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "jusene"
	user3.Gender = "man"
	user3.Address.CreateTime = time.Now()
	user3.Email.CreateTime = time.Now()

	fmt.Printf("%#v", user3)
}
