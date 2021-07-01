package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Card struct {
	ID    int64   `json:"id,string"` // 添加string tag
	Score float64 `json:"score,string"`
}

func main() {
	jsonStr1 := `{"id": "123456", "score": "98.5"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		log.Fatalf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
	}
	fmt.Printf("c1:%#v\n", c1)
}
