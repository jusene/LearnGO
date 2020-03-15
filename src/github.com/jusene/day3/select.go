package main

import "fmt"

func main() {
	a := make(chan int, 1024) //创建一个channel
	b := make(chan int, 1024)

	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次，", i)
		a <- 1
		b <- 1
		select {
		case <-a:
			fmt.Println("from a")
		case <-b:
			fmt.Println("from b")
		}
	}
}
