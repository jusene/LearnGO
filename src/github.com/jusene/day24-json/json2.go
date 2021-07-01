package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Hobby []string `json:"hobby"`
}

// 在tag中添加omitempty忽略空值
type User2 struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}

func main() {
	u1 := User1{
		Name: "jusene",
	}

	u2 := User2{
		Name: "jusene",
	}

	b1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	b2, err := json.Marshal(u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))
	fmt.Println(string(b2))
}
