package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	StuID int `json:"id"` // 通过指定tag实现json序列化该字段时的key
	Gender string
	name string
}

func main() {
	s1 := Student{
		StuID:     1,
		Gender: "man",
		name:   "jusene",
	}

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Println(string(data)) // {"id":1,"Gender":"man"} name 小写字母开头，小写表示私有，只能在结构体中使用
}

