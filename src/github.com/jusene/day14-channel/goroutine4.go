package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 表示将cpu的时间让给别人
		fmt.Println(i, s)
	}
}

func main() {
	go say("golang")
	say("hello")
}
