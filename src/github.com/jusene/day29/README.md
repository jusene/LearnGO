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

- Background() 主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context
- TODO() 不知道使用什么Context的时候，可以使用这个

### WithCancel

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

```go
package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <- chan int {
	det := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <- ctx.Done():
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
```

### WithDeadline

```
func WithDeadline(parent Context, deadline time.time) (Context, CancelFunc)
```

```go
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
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

### WithTimeout

```
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。具体示例如下：

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
	LOOP:
		for {
			fmt.Println("db connecting ...")
			time.Sleep(time.Millisecond * 10) //假设正常连接数据库耗时
			select {
			case <- ctx.Done(): // 50 毫秒后自动调用
			break LOOP
			default:

			}
		}
		fmt.Println("worker done")
		wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知goroutine结束
	wg.Wait()
	fmt.Println("over")
}
```

### WithValue

```
func WithValue(parent Context, key, val interface{}) Context
```

WithValue返回父节点的副本，其中与key关联的值为val。

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。

```go
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
			case <- ctx.Done():
				break LOOP
			default:

			}
		}
		fmt.Println("worker done")
		wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "21212121")

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
```

## 使用Context的注意事项

- 推荐以参数的方式显示传递Context
- 以Context作为参数的函数方法，应该把Context作为第一个参数
- 给一个函数方法传递Context的时候，不要传递nil，不知道就传context.TODO()
- Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
- Context是线程安全的，可以放心在多个goroutine中传递

