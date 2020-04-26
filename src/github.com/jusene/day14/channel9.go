package main

import "fmt"

func counter(out chan <- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan <- int, in <- chan int) { // out单向通道，只能发送不能接收；in单向通道，只能接收不能发送
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <- chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
