## Go标准库fmt

### Print

`Print`系列函数会将内容输出到系统的标准输出。

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

```go
package main

import "fmt"

func main() {
	fmt.Print("输出1")
	name := "jusene"
	fmt.Printf("%s\n", name)
	fmt.Println("hello world")
}
```

### Fprint

Fprint系列函数将内容输出到一个`io.Writer`接口类型的变量`w`中。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Weiter, a ...interface{}) (n int, err error)
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "标准输出")
	fmt.Fprint(os.Stderr, "标准错误输出")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错， err:", err)
		return
	}

	name := "jusene"
	fmt.Fprintf(fileObj, "往文件写入信息: %s", name)
}
```
`只要满足io.Writer接口的类型都支持写入`

### Sprint

`Sprint`系列函数会把传入的数据生成一个字符串。

```go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

```go
package main

import "fmt"

func main() {
	name := fmt.Sprint("jusene")
	age := 27
	n := fmt.Sprintf("name:%s age:%d", name, age)
	m := fmt.Sprintln("坚持")
	fmt.Println(n, m)
}
```

### Errorf

`Errorf`函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
		}
	}()

	age, err := funErr()
	if err != nil {
		panic(err)
	}
	fmt.Print(age)
}

func funErr() (age int,err error) {
	return 100, fmt.Errorf("错误")
}
```

```go
e := error.New("错误e")
w := fmt.Errorf("Wrap了一个错误%w", e)
```
`%w占位符用来生成一个可以包裹Error的Wrapping Error`

### 格式化占位符

- 通用符

|  占位符   | 说明  |
|  ----  | ----  |
| %v  | 值的默认格式 |
| %+v | 类似%v，但输出结构体时会添加字段名 |
| %#v | 值的Go语法表示 |
| %T | 打印值的类型 |
| %% | 百分号 |

```go
package main

import (
	"fmt"
)

func main() {
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%")
}

{小王子}
{name:小王子}
struct { name string }{name:"小王子"}
struct { name string }
100%
```

- 布尔型

|  占位符   | 说明  |
|  ----  | ----  |
| %t | true或false |

- 整型

|  占位符   | 说明  |
|  ----  | ----  |
| %b | 表示二进制 |
| %c | 该值对应的unicode码值 |
| %d | 表示10进制 |
| %o | 表示八进制 |
| %x | 表示十六进制,a-f |
| %X | 表示十六进制,A-F |
| %U | 表示为Unicode格式 |
| %q | 该值对应的单引号括起来的go语法字符字面量 |

- 浮点数与复数

|  占位符   | 说明  |
|  ----  | ----  |
| %b | 无小数部分、二进制指数的科学计数法 |
| %e | 科学计数法，如-1234.456e+78 |
| %E | 科学计数法，如-1234.456E+78 |
| %f | 有小数部分但无指数部分，如123.456 |
| %F | 等价于%f |
| %g | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
| %G | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

- 字符串和[]byte

|  占位符   | 说明  |
|  ----  | ----  |
| %s | 直接输出字符串或者[]byte |
| %q | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x | 每个字节用两字符十六进制数表示（使用a-f |
| %X | 每个字节用两字符十六进制数表示（使用A-F）|

- 指针

|  占位符   | 说明  |
|  ----  | ----  |
| %p | 表示为十六进制，并加上前导的0x |

### 宽度标识符

|  占位符   | 说明  |
|  ----  | ----  |
| %f | 默认宽度，默认精度 |
| %9f | 宽度9，默认精度 |
| %.2f | 默认宽度，精度2 |
| %9.2f | 宽度9，精度2 |
| %9.f | 宽度9，精度0 |

```go
package main

import "fmt"

func main() {
	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)
}

12.340000
12.340000
12.34
    12.34
       12
```

### 其他flag

|  占位符   | 说明  |
|  ----  | ----  |
| '+' | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）|
| '' | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
| '-' | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）|
| '#' | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值 |
| 'o' | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面 |

```go
import "fmt"

func main() {
	s := "你好"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}
```

### 捕获输入

Go语言包下`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，可以在程序运行过程中从标准输入获取用户的输入。

```
func Scan(a ...interface{}) {n int, err error}
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
```

```go
package main

import "fmt"

func main() {
	var (
		name string
		age int
		married bool
	)
	fmt.Print("请输入用户:")
	fmt.Scan(&name) // 读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
	fmt.Println()
	fmt.Print("请输入年龄:")
	fmt.Scanln(&age) // 遇到换行时才停止扫描
	fmt.Print("请输入婚姻:")
	fmt.Scanf("%t", &married) // 读取由空白符分隔的值保存到传递给本函数的参数中
	fmt.Println()
	
	fmt.Printf("name:%s age:%d married:%t \n", name, age, married)
```

- bufio.NewReader

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bufioDemo()
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
```

- Fscan

Fscan不是从标准输入中读取数据而是从`io.Reader`中读取数据

```
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
```

- Sscan

Sscan不是从标准输入中读取数据而是从指定字符串中读取数据

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```