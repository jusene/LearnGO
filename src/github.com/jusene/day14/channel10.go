package main

import (
	"fmt"
)

func send(ch chan int)  {
	ch <- 10
	close(ch)
}

func main() {
	ch := make(chan int)
	go send(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
