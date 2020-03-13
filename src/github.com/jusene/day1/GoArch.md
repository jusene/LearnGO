## Go程序元素

### 标识符

- 用户定义标识符，作为程序实体存在，一般来说同一个代码块中不允许程序实体同名,使用不同代码包中的实体需要使用包标识符。
- 预定义标识符，数据类型，内建函数名等

### 关键字

数据类型：
```
bool bytes float32 float64 
int int8 int16 uint32 int64 
uint uint8 uint16 uint32 uint64
string uint uint8 uintptr complex128
complex64 rune error
```

内建函数名：
```
beak     default      func      interface     select
case     defer        go        map           struct
chan     else         goto      package       switch
const    fallthrough  if        range         type
continue  for         import    return        var
append    make        len       cap           new
copy      delete      close     complex       real
imag      panic       recover   
```

常量：
```
true false iota  nil
```

### 字面量

字面量就是表示值的一种标记法，首先用于表示基础数据类型值的各种字面量，其次表示用户构造的自定义复合数据类型的类型字面量，最后它还是复合数据类型的值的复合字面量。

``` 
const b int = 10 // b为常量，10为字面量
var str string = "hello world" // str为变量，hello world为字面量
```

### 运算符

二元运算符优先级

```
* / % << >> & &^  最高
+ = | ^ 较高
== != < <= > >= 中
&& 较低
|| 最低
```

## Go语言基本概念

### 常量

常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。

``` 
const pi = 3.1415
const e = 2.7182
```

多个常量一起声明
``` 
const (
    pi = 3.1415
    e = 2.7182
)
```

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
``` 
const (
    n1 = 100 // 100
    n2       // 100
    n3       // 100
)
```

### 枚举

iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。 

```go
package main

import "fmt"

const (
    a = iota // 0
    b        // 1
    c        // 2
    d, e, f = iota, iota, iota // 3 3 3
    g = iota // 4
    h = "h"  // h 但是iota累加
    i        // 默认使用上面的赋值，iota依然累加
    j = iota // 7
    _ = iota // 累加
    k        // 9
)

const z = iota

func main() {
    fmt.Println(a, b, c, d, e, f, g, h, i, j, k, z)
}

0 1 2 3 3 3 4 h h 7 9 0
```

定义数量级
``` 
const (
		_  = iota
		KB = 1 << (10 * iota) // 二进制左移10位 2^10=1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```

### 变量

常见变量的数据类型有：整型、浮点型、布尔型等

``` 
var (
        a string
        b int
        c bool
        d float32
    )
```

变量的初始值：
- 整型和浮点型 0
- 字符串 空字符串
- 布尔型 false
- 切片，函数，指针变量 nil

一次初始化多个变量
``` 
var name, age = "jusene", 26
```

类型推导
``` 
var name = "jusene"
var age = 26
```

短变量声明
``` 
a := 1
```
官方推荐，但是只能使用在函数体内

匿名变量
```go
packge main

import (
    "fmt"
)

func GetClass() (stuNum int, className, headTeacher string) {
    return 49, "一班", "张三"
}

func main() {
    stuNum, _, _ := GetClass()
    fmt.Println(stuNum)
}
```

### init函数

可以在init函数中完成初始化，init函数是一个特殊的函数，优先级比main函数高，并且不能手动调用init函数，每一个源文件有且只能有一个init函数
```go
package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

func main() {
	DPi := Pi * Pi
	fmt.Println(Pi, DPi)
}
```

