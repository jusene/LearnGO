package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	p1 := Person{
		Name:   "jusene",
		Age:    18,
		Weight: 71.5,
	}

	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		log.Fatalf("json marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%s\n", b)

	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		log.Fatalf("json unmarshal failed, err:%v\n", err)
	}
	fmt.Printf("p2:%v\n", p2)
}
