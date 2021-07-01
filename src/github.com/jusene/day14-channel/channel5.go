package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	fmt.Println(len(ch), cap(ch))
	ch <- 10 // 有缓冲的通道不会死锁
	fmt.Println(len(ch), cap(ch))
	fmt.Println("发送成功")
	time.Sleep(2 * time.Second)
	value := <- ch
	fmt.Println("接收成功", value)
}
