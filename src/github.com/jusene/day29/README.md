## GO语言 Context

### Context

Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

example:
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:
		}
	}
}

func worker2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go worker(ctx)
	go worker2(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
```

### Context接口

定义4个具体方法

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）
- Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel
- Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值：
- - 如果当前Context被取消就会返回Canceled错误
- - 如果当前Context超时就会返回DeadlineExceeded错误
- Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据

### Background()和TODO()

Go内置的两个函数: Background()和TODO()

