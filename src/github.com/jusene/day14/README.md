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

### select多路复用

select类似switch语句，每个case会对应一个通信的通道（接收或发送），select会一直等待，直到某个case的通信操作完成时。select有较多的限制，其中最大的一条限制就是每个case语句必须有一个IO操作

```
select {
    case ch1:
        ...
    case ch2:
        ...
    default:
        ...
}
```

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <- ch: // 如果chan1成功读取到数据，则进行该case处理语句
			fmt.Println(x) 
		case ch <- i: // 如果成功向chan2写入数据，则进行该case处理语句
        default: // 如果上面都没成功，则进入default处理流程
        }
	}
}
```

- 可以处理一个或多个channel的发送和接受操作
- 如果多个case满足，select会随机选择一个
- 对于没有case的select{}会一直等待，可用于阻塞main函数

### 并发安全和锁

```go
package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

以上的并发导致数据竞争，每次结果不同

#### 互斥锁

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。

```go
package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

#### 读写互斥锁

当需要并发读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景可以使用读写锁。

读写锁分为：读锁和写锁，当一个goroutine获取读锁的时候，其他goroutine如果获取读锁会继续获得锁，如果获得写锁就会等待；当一个goroutine获取写锁后，其他goroutine无论是获取读锁还是写锁都会等待。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock() // 加写锁
	time.Sleep(time.Millisecond)
	rwlock.Unlock() // 解写锁
	wg.Done()
}

func read() {
	rwlock.RLock()  // 加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()

	// 写锁并发时间18.9736ms
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}


	// 读锁并发时间1.0203ms
	/*
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	*/
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
```

读写锁非常适合读多写少的场景

#### sync.WaitGroup

| 方法名 | 功能 |
| ----- | ----- |
| (wg *WaitGroup) Add(delta int) | 计数器+delta |
| (wg *WaitGroup) Done() | 计时器-1 |
| (wg *WaitGroup) Wait() | 阻塞直到计数器变为0 |

#### sync.Once

确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

```
func (o *once) Do(f func()) {}
```

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

func main() {

	for i, v := range make([]string, 10) {
		//once.Do(onces) // once只会被执行一次
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(onced)
			fmt.Println("213")
			wg.Done()
		}()
	}
	wg.Wait()
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
```

并发安全的单例模式

```go
import (
	"fmt"
	"sync"
	"time"
)

type singleton struct {
	name string
	age int
}

var instance *singleton
var once sync.Once
var wg sync.WaitGroup
var lock sync.Mutex

func GetInstance() *singleton {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("load instance...")
		instance = &singleton{
			name: "jusene",
			age: 27,
		}
		fmt.Println(instance.name, instance.age)
	})
	lock.Lock()
	fmt.Println("*** 过了一年")
	time.Sleep(time.Second)
	instance.age += 1
	lock.Unlock()
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetInstance()
	}
	wg.Wait()

	fmt.Println(instance.name, instance.age)
}
```

#### sync.Map

Go语言中内置的map不是并发安全的，需要使用sync.Map

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i ++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.LoadOrStore(key, 100)
			fmt.Printf("k=%v, v=%v\n", key, value)
		}(i)
	}
	wg.Wait()
}
```