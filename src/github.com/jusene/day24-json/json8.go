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
	//jsonStr := `{"create_time":"2020-07-10T09:57:46.0743201+08:00"}`
	jsonStr := `{"create_time":"2020-07-10 09:57:46"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshl failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
