package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) //假设正常连接数据库耗时
		select {
		case <-ctx.Done(): // 50 毫秒后自动调用
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知goroutine结束
	wg.Wait()
	fmt.Println("over")
}
