package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User              //匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func main() {
	u1 := User{
		Name:     "jusene",
		Password: "1234567",
	}

	b, err := json.Marshal(PublicUser{
		User: &u1,
	})
	if err != nil {
		log.Fatalf("json.Marshal u1 failed, err:%v\n", err)
	}

	fmt.Printf("str:%s\n", b)
}
