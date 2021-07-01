package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	det := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case det <- n:
				n++
			}
		}
	}()
	return det
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
