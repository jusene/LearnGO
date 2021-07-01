package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Blog    string `json:"blog"`
}

func main() {
	u1 := User{
		Name: "jusene",
	}

	b1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))
}
