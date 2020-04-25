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

```


