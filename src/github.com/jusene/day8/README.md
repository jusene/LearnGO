## Go语言函数

```
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    // 逻辑代码
    return value1, value2 // 返回多值
}
```

函数名称如果小写开头，它的作用只属于所证明的包，不能被其他包调用；如果函数名以大写字母开头，则该函数便是公开的，可以被其他包调用。

Go语言函数不支持嵌套，重载和默认参数。

```go
package main

import "fmt"

func main() {
	x := 3
	y := 4
	a, b := SumAndProduct(x, y)
	fmt.Println(a, b)
	c, d := SumAndProduct1(x, y)
	fmt.Println(c, d)
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

// 命名的返回值就相当于在函数的声明的时候声明了一个变量
func SumAndProduct1(A, B int) (add int, mul int) {
	add = A + B
	mul = A * B
	return
}
```

### 函数作为参数

```go
package main

import "fmt"

func main() {
	num := 100
	filter(isBig, num)
}

type funcBool func(x int) bool

func filter(f funcBool, num int) bool {
	if f(num) {
		fmt.Println("ok")
		return true
	} else {
		fmt.Println("no")
		return false
	}
}

func isBig(x int) bool {
	if x >= 100 {
		return true
	}
	return false
}
```

### 可变参数

```go
package main

import (
	"fmt"
)

func main() {
	age := ageMinOrMax("min", 1, 3, 4, 59)
	fmt.Printf("年龄最小的参数是%d岁\n", age)

	ageArr := []int{6, 4, 43, 2, 32}
	age = ageMinOrMax("max", ageArr...)
	fmt.Printf("年龄最大的参数是%d岁\n", age)
}

func ageMinOrMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}

	if m == "max" {
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	} else if m == "min" {
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	} else {
		e := -1
		return e
	}
}
```

### 返回值

```go
package  main

import "fmt"

func main()  {
	ret := dosome()
	if ret == nil {
		fmt.Println(ret)
	}

}

func dosome() []int {
	return nil // nil可以看作一个有效的slice, 没必要显示返回一个长度为0的切片[]int{}
}
```

### defer语句

defer语句延迟处理，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

```go
package main

import "fmt"

func main() {
	fmt.Println("打开文件")
	defer fmt.Println("关闭远程连接")
	defer fmt.Println("关闭文件")
	fmt.Println("读取文件")
}
```

### 变量作用域

```go
package main

import "fmt"

// 定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Println(num)
}

func main() {
	testNum() // 当局部变量和全局变量重名，优先使用局部变量
}
```

for, if, switch语句块使用这种变量也只可以在当前块中使用

```go
package main

import "fmt"

func main() {
	testLocalVars(20, 2)
}

func testLocalVars(x, y int) {
	fmt.Println(x, y)
	if x > 10 {
		z := 100 // 变量z只在if块中生效
		fmt.Println(z)
	}
	// fmt.Println(z) 无法使用变量z
}
```

### 函数类型与变量

```go
package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func main() {
	var c calculation // 申明一个calculation类型的变量
	c = add
	fmt.Printf("%T", c) // main.calculation
	fmt.Println(c(1, 3))

	f := add          // 将add函数赋值给变量f
	fmt.Printf("%T", f) // func(int, int) int
	fmt.Println(f(1, 3))
}
```

### 高阶函数

函数作为参数

```go
package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	ret := calc(10, 20, add)
	fmt.Println(ret)
}
```

函数作为返回值

```go
package main

import (
	"errors"
	"fmt"
)

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "*":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x * y
}

func main() {
	op, err	:= do("+")
	if err != nil {
		panic(err)
	}
	ret := op(10, 20)
	fmt.Println(ret)
}
```

### 匿名函数与闭包

```
func(参数)(返回值) {
    函数体
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
package main

import "fmt"

func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) int {
		return x + y
	}
	ret := add(10, 20)
	fmt.Println(ret)

	// 自执行函数，匿名函数定义加()直接完成
	ret1 := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println(ret1)
}
```

匿名函数多用于实现回调函数和闭包

### 闭包

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。

```go
package main

import "fmt"

func Add() func(b int) int {
	var a int
	return func(b int) int {
		a += b
		return a
	}
}

func main() {
	var f = Add()
	fmt.Println(f(10)) // 10
	fmt.Println(f(10)) // 20
	fmt.Println(f(10)) // 30
}
```

整个函数的生命周期中应用了外部作用域a的变量，此时函数f就是一个闭包，在f的生命周期中a是一直有效的。

```go
package main

import "fmt"

func Add(a int) func(b int) int {
	return func(b int) int {
		a += b
		return a
	}
}

func main() {
	var f = Add(10)
	fmt.Println(f(10)) // 10
	fmt.Println(f(10)) // 20
	fmt.Println(f(10)) // 30
}
```

闭包进阶：
```go
package main

import (
	"fmt"
	"strings"
)

func checkFile(filename string) func(prefix, suffix string) bool {
	return func(prefix, suffix string) bool {
		if strings.HasPrefix(filename, prefix) && strings.HasSuffix(filename, suffix) {
			return true
		}
		return false
	}
}

func main() {
	jpgFunc := checkFile("test.jpg")
	fmt.Println(jpgFunc("test", "jpg"))
	fmt.Println(jpgFunc("test", "txt"))
}
```

闭包进阶：
```go
package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(1), f2(2)) // 10 8
}
```

闭包进阶:
```go

```