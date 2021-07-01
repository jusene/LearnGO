package main

import "fmt"

func count(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	count(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
