package main

import "fmt"

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)  // 无缓冲的通道只有在有人接受值的时候才能发送
	ch <- 1 // deadlock 死锁
	fmt.Println("发送成功")
}
