package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	select {
	case ch <- 0:
		v := <- ch
		fmt.Println("接受到ch", v)
	case ch <- 1:
		v := <- ch
		fmt.Println("接受到ch", v)
	}
}
