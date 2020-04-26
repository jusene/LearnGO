package main

import (
	"fmt"
	"time"
)

func count(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func rec(ch chan int) {
	for {
		i, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

func rec1(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go count(ch)
	//go rec(ch)
	go rec1(ch)
	time.Sleep(time.Second)
}
