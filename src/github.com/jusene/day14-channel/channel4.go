package main

import (
	"fmt"
)

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	ch <- 10 // 死锁
	fmt.Println("发送成功")

	value := <- ch // 无缓冲的通道只能在有人接收值得时候才能发消息，所以只能另起goroutine来接收
	fmt.Println(value)
}
