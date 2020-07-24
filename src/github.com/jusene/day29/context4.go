package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond) // 过期
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// ctx会过期，任何情况下调用它的cancel函数都是很好的实践
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
