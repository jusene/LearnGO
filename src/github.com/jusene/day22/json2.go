package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type MyDATA struct {
	Name  string  `json:"item"`
	Other float64 `json:"amount"`
}

func main() {
	s := new(MyDATA)
	s.Name = "jusene"
	s.Other = 0.23
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	js, err := simplejson.NewJson(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(js.Get("item").String())
}
