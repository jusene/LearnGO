package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	u1 := UserInfo{
		ID:   123456,
		Name: "JUSENE",
	}

	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"1212121212",
	})
	if err != nil {
		fmt.Printf("json.Marsha failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}
