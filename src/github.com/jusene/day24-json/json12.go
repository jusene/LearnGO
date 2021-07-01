package main

import (
	"encoding/json"
	"fmt"
)

type sendMsg struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func main() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`

	// 定义一个map，value类型为json.RawMessage，方便后续更灵活地处理
	var data map[string]json.RawMessage

	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	fmt.Printf("data:%#v\n", data)
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)
}
