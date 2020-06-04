package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	CreateTime time.Time `json:"create_time"`
}

func main() {
	p1 := Post{CreateTime: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-06-04 13:28:26.3052599+08:00"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshl failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
