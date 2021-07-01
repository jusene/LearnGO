package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	TraceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invaild trace code")
	}

LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", TraceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "21212121")

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
