## Go语言流程控制

Go语言中特殊的关键字：
- defer: 用于捕获异常和资源回收
- select: 用于分支选择（配合通道使用）
- go: 用于异步启动goroutine并启动特定函数
- goto: 跳转到指定的label

### 条件语句

if条件格式:
```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

```go
package main

import "fmt"

func main() {
    a := 13
    if a > 20 {
        fmt.Printf("a大于20\n")
    } else if a < 10 {
        fmt.Printf("a小于10\n")
    } else if a == 11 {
        fmt.Printf("a等于11\n")
    } else {
        fmt.Printf("a大于10\n")
        fmt.Printf("a小于20\n")
        fmt.Printf("a不等于11\n")
    }
}
```

初始化子语句

```go
package main

import "fmt"

func main() {
    if a := 10; a < 20 {
        fmt.Printf("a小于20\n")
    } else {
        fmt.Printf("a的值是: %d\n", a)
    }
}
```

- 子语句只能有一个表达式
- a的值在if代码块中定义，所有不能在代码块之外调用

### 循环语句

在Go语言中，循环语句中关键字是for，没有while关键字

for 循环的基本格式:
```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

```go
package main

import "fmt"

func main() {
    for a := 1;a < 5;a++ {
        fmt.Printf("a的值是：%d\n", a)
    }
}
```

for 语句的三个子语句：初始化子语句、条件子语句、后置子语句，其中条件子语句是必须的。

```go
package main

import "fmt"

func main() {
    a := 0
    b := 5
    for a < b {
        a++
        fmt.Printf("a的值是：%d\n", a)
    }
}
```

***无限循环***

```go
package main

import "fmt"

func main() {
    a := 0
    b := 5
    for a < b {  // 当条件子语句一直满足的时候，会一直循环
        fmt.Printf("a的值是：%d\n", a)
    }
}
```

Go语言有默认的无限循环结构

```go
for {
    循环体语句
}
```
for循环可以通过break、goto、return、panic语句强制退出循环。

***for range子语句***

每一个for都可以使用一个特殊的range子语句，其作用类似于迭代器，用于轮询数组或者切片值中的每一个元素，也可以用于轮询字符串的每一个字符，以及字典值中的每一个键值对，甚至还可以持续读取每一个通道类型中的元素。

- 数组、切片、字符串返回索引和值。
- map返回键和值。
- 通道（channel）只返回通道内的值。

```go
package main

import "fmt"

func main() {
	str := "jusene"
	for i, char := range str {
		fmt.Printf("字符串第%d个字符串的值为%d\n",i, char)
	}
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
	numbers := []int{1, 2, 3}
	for i, x := range numbers {
		fmt.Println(i, x)
	}
	nums := [5]int{1,2,3,4}
	for i, x := range nums {
	    fmt.Println(i, x)   
	}
}
```

### 选择语句

#### 表达式switch

```go
package main

import "fmt"

func main() {
    var grade string 
    marks := 90
    
    switch marks {
        case 90:
            grade = "A"
        case 80:
            grade = "B"
        case 60, 70:
            grade = "C"
        default:
            grade = "D"
    }
    
    fmt.Printf("你的成绩为%s\n", grade)
}
```

switch 初始化语句

```go
package main

import "fmt"

func main() {
	var grade string
	switch marks := 90; marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 60:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("你的成绩为%s\n", grade)
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量

```go
package main

import "fmt"

func main() {
	var grade string
	marks := 88

	switch  {
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch  {
	case grade == "A":
		fmt.Println("成绩优秀！")
	case grade == "B":
		fmt.Println("表现良好！")
	case grade == "C", grade == "D" :
		fmt.Println("再接再厉！")
	default:
		fmt.Println("成绩不合格！")
	}

	fmt.Printf("你的成绩为%s\n", grade)
}
```

switch初始化语句

```go
package main

import "fmt"

func main() {
	var grade string  
    switch marks := 90; { // 分号不能省略
        case marks >= 90:
            grade = "A"
        case marks >= 80:
            grade = "B"
        case marks >= 70:
            grade = "C"
        case marks >= 60:
            grade = "D"
        default:
            grade = "E"
    }
    fmt.Printf("你的成绩为%s\n", grade)
}
```

fallthrough关键词可以把当前case控制权交给下一个case语句判断

```go
package main

import "fmt"

func main() {
	var grade string  
    switch marks := 90; { // 分号不能省略
        case marks >= 90:
            fallthrough
        case marks >= 80:
            grade = "B"
        case marks >= 70:
            grade = "C"
        case marks >= 60:
            grade = "D"
        default:
            grade = "E"
    }
    fmt.Printf("你的成绩为%s\n", grade)
}
```

类型switch

```go
package main

import "fmt"

var x interface{}

func main() {
	x = 1
	switch i := x.(type) {
	case nil:
		fmt.Printf("这里是nil，x的类型是%T", i)
	case int:
		fmt.Printf("这里是int，x的类型是%T", i)
	case float64:
		fmt.Printf("这里是float64，x的类型是%T", i)
	case bool:
		fmt.Printf("这里是bool，x的类型是%T", i)
	case string:
		fmt.Printf("这里是string，x的类型是%T", i)
	default:
		fmt.Printf("未知类型")
	}
}
```

#### 表达式select

switch按照顺序从上到下执行，而select是随机选择一个case来判断，select用于配合通道channel的读写操作，用于多个channel的并发读写操作。

```go
package main

import "fmt"

func main() {
    a := make(chan int, 1024) //创建一个channel
    b := make(chan int, 1024)
    
    for i := 0; i < 10; i++ {
        fmt.Printf("第%d次，", i)
        a <- 1
        b <- 1
        select {
            case <-a:
                fmt.Println("from a")
            case <-b:
                fmt.Println("from b")
        }
    }
}
```

### 标签

在go语言中，有一种特殊概念就是标签，可以给for，switch，select等流程控制代码打一个标签，配合标签标识符可以方便跳转到某一个地方继续执行。

```go
L1:
    for i := 0;i <= 5;i++ {
        //代码块
    }
    
L2:
    switch i {
        //代码块
    }
```

#### goto(跳转到指定标签)

跳出两层循环

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 2 {
				// 跳到标签
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}
	
	breakTag:
		fmt.Println("结束两层for循环")
}
```

goto只能在同一个函数内跳转

#### break(跳出循环)

break会跳出当前循环，如何跳出指定代码块，需要结合标签进行

```go
package main 

import "fmt"

func main() {
    for {
        x := 1
        switch {
            case x > 0:
            fmt.Println("A") // 无限循环
            break
            case x == 1:
            fmt.Println("B")
            default:
            fmt.Println("C")
        }
    }
}
```

```go
package main 

import "fmt"

func main() {
LOOP1:
    for {
        x := 1
        switch {
            case x > 0:
            fmt.Println("A")
            break LOOP1
            case x == 1:
            fmt.Println("B")
            default:
            fmt.Println("C")
        }
    }
}
```

#### continue(继续下次循环)

ontinue语句可以结束当前循环，开始下一次的循环迭代过程

```go
package main

import "fmt"

func main() {
    LOOP1:
    for i := 0;i <= 5;i++ {
        switch {
            case i > 0:
            fmt.Println("A")
            continue LOOP1
            case i == 1:
            fmt.Println("B")
            default:
            fmt.Println("C")
        }
        fmt.Printf("i is: %d\n", i)
    }
}
```

### 延迟语句

defer就是其中一个，defer用于延迟调用指定函数，defer关键字只能出现在函数内部，defer语句会在函数最后执行

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Printf("hello ")
}
```

defer有两大特点:

- 只有当defer语句全部执行，defer所在的函数才算真正结束执行
- 当函数中有defer语句时，需要等待所有defer语句执行完毕，才会执行return语句

defer语句用于回收资源、清楚收尾等工作

```go
package main 

import "fmt"

var i = 0

func print() {
    fmt.Println(i, &i)
}

func main() {
    for ; i < 5; i++ {
        defer print()
    }
}

5 0x11a5b00
5 0x11a5b00
5 0x11a5b00
5 0x11a5b00
5 0x11a5b00
```

```go
package main

import "fmt"

var i = 0

func print(i int) {
    fmt.Println(i, &i)
}

func main() {
    for ; i < 5; i++ {
        defer print(i)
    }
}

4 0xc0000b4008
3 0xc0000b4030
2 0xc0000b4040
1 0xc0000b4050
0 0xc0000b4060
```
