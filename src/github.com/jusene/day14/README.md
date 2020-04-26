## Go语言并发

Go语言的并发通过goroutine实现，goroutine是由go运行时runtime调度完成，而线程由操作系统调度完成。

Go语言还提供channel在多个goroutine间进行通信。

### 使用goroutine

Go语言中使用goroutine，只要在调用函数的时候在前面加上go关键字

```go
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine")
}

func main() {
	go hello()
	fmt.Println("main goroutine done!")
	time.Sleep(2 * time.Second)
}
```

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int)  {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Printf("Hello Goroutine%d\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所以登记的goroutine都结束
	fmt.Println("all goroutine finished")
}
```

这些10个goroutine是并发执行的，而goroutine的调度是随机的。

### GOMAXPROCS

Go运行的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码，默认是机器上的CPU核心数。

Go语言通过`runtime.GOMAXPROCS()`函数设置当前程序并发占用的CPU逻辑核心数。

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A: ", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B: ", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
```

Go语言中的操作系统线程和goroutine的关系：
- 一个操作系统线程对应用户态多个goroutine
- go程序可以同时使用多个操作系统线程
- goroutine和OS线程是多对多的关系

### channel

#### channel类型

channel是一种类型，一种引用类型

```
var 变量 chan 元素类型
```

```
var ch1 chan int // 声明一个传递整数的通道
var ch2 chan bool // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
```

#### 创建channel

通道是引用类型，通道类型的空值是nil

```
var ch chan int
fmt.Println(ch) // nil
```

声明的通道需要使用make函数初始化之后才能使用

```
make(chan 元素, [缓冲大小])
```

#### channel操作

通道有发送(send)、接受(receive)、关闭(close)三种操作

发送和接受都使用`<-`符号

关闭通道需要注意，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。

- 对一个关闭的通道再发送就会导致panic
- 对一个关闭的通道进行接受会一直获取值直到通道为空
- 对一个关闭的并且没有值得通道执行接受操作会得到对应类型的零值
- 关闭一个已经关闭的通道会导致panic

```go
package main

import (
	"fmt"
	"time"
)

func Count(ch chan int, i int) {
	ch <- i
	fmt.Println("Counting")
}

func main() {
	chs := make([] chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int, 10)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		value := <- ch
		fmt.Println(value)
	}
	time.Sleep(5 * time.Second)
}
```

#### 无缓冲的通道

无缓冲通道称为阻塞通道

```go
package main

import (
	"fmt"
)

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	ch <- 10 // 死锁
	fmt.Println("发送成功")

	value := <- ch // 无缓冲的通道只能在有人接收值得时候才能发消息，所以只能另起goroutine来接收
	fmt.Println(value)
}
```

使用无缓冲通道进行通信将导致发送和接收的goroutine同步化，无缓冲通道也被称为同步通道。

#### 有缓冲的通道

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	fmt.Println(len(ch), cap(ch))
	ch <- 10 // 有缓冲的通道不会死锁
	fmt.Println(len(ch), cap(ch))
	fmt.Println("发送成功")
	time.Sleep(2 * time.Second)
	value := <- ch
	fmt.Println("接收成功", value)
}
```

#### 单向通道

在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收

```go
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
```

#### 通道总结

| channel | nil | 非空 | 空的 | 满了 | 没满 | 
| ----- | ----- | ----- | ----- | ----- | ----- |
| 接收 | 阻塞 | 接收值 | 阻塞 | 接收值 | 接收值 |
| 发送 | 阻塞 | 发送值 | 发送值 | 阻塞 | 发送值 |
| 关闭 | panic | 关闭成功，读完数据后返回零值 | 关闭成功，返回零值 | 关闭成功，读完数据后返回零值 | 关闭成功，读完数据后返回零值 |

关闭已经关闭的channel也会引发panic

### goroutine池

work pool，防止goroutine泄露和暴涨

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <- chan int, result chan <- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 输出结果
	for a := 1; a <= 5; a++ {
		<- results
	}
}
```