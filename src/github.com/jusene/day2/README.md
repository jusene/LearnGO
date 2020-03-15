## Go基础数据类型

Go语言基本的数据类型：整型、浮点型、布尔型、字符串，还有数组、切片、结构体、函数、map、通道（channel）等。

### 整型

|  类型   | 描述  |
|  ----  | ----  |
| int8  | -128~127 |
| uint8  | 0~255 |
| int16 | -32768~32767 |
| uint16 | 0~65535 |
| int32 | -2147483648~2147483647 |
| uint32 | 0~4294967295 |
| int64 | -9223372036854775808~9223372036854775807 |
| uint64 | 0~18446744073709551615 |

特殊整型

| 类型 | 描述 |
| ---- | ---- |
| uint | 32位操作系统上就是uint32，64位操作系统上就是uint64 |
| int | 32位操作系统上就是int32，64位操作系统上就是int64 |
| uintptr | 无符号整型，用于存放一个指针 |

存在操作系统的差异，存在不可控因素，不要用int和uint，当然虽然都是整型，但是不同的类型还是不可比较的。

```go
package main

import "fmt"

func main() {
	// 十进制
	var a int64 = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 二进制 1010
	fmt.Printf("%o \n", a) // 八进制 12
	fmt.Printf("%x \n", a) // 十六进制 a

	// 八进制 以0开头
	var b int64 = 077
	fmt.Printf("%d \n", b) // 十进制 63
	fmt.Printf("%b \n", b) // 二进制 111111
	fmt.Printf("%o \n", b) // 77
	fmt.Printf("%x \n", b) // 十六进制 3f

	// 十六进制 以0x开始
	var c int64 = 0xff
	fmt.Printf("%d \n", c) // 十进制 255
	fmt.Printf("%b \n", c) // 二进制 11111111
	fmt.Printf("%o \n", c) // 八进制 377
	fmt.Printf("%x \n", c) // ff
}
```

## 浮点型

float32浮点数的最大范围3.4e38，可以使用定义math.MaxFloat32；float64的浮点数的最大范围1.8e308，可以使用定义math.MaxFloat64。

```go
package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a := 3.0
	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi)
	fmt.Println(reflect.TypeOf(a)) // 默认类型float64
}
```

### 复数

- complex64 // 32位实数和虚数
- complex128 // 64位实数和虚数

```go
package main

import "fmt"

func main() {
    var v1 complex64
    v1 = 3.2 + 12i
    v2 := 2.3 + 12i
    v3 := complex(3.2, 12)
    v := v2 + v3
    fmt.Println(v1, v2, v3, v)
}
```

### 布尔型

- 布尔类型变量的默认值为false。
- Go 语言中不允许将整型强制转换为布尔型.
- 布尔型无法参与数值运算，也无法与其他类型进行转换。

### 字符串

Go语言中单引号代表字符，双引号代表字符串，反引号代表非解释字符串

`字符串是不可修改的`

***Go支持两种形式的字面量***

> 解释字符串
- \a 响铃
- \b 回车
- \f 换页
- \n 换行
- \r 回车
- \t 制表符
- \u unicode字符
- \v 垂直制表符
- \\' 单引号 （只用在'形式的rune符号面值中）
- \\" 双引号 （只用在"形式的rune符号面值中）
- \\  反斜杠

> 非解释字符串
```go
package main

import "fmt"

func main() {
    str1 := `苟利国家生死以\n岂因祸福避趋之`
    str2 := "苟利国家生死以\n岂因祸福避趋之"
        
    fmt.Println(str1, str2)
}
```

#### 字符串遍历

```go
package main

import "fmt"

var str string = "武汉加油"

func main()  {
	for i := 0; i < len(str); i++ {
		fmt.Println(i)
		fmt.Printf("%c", str[i])
	} // 输出字节码

	for i, v := range str {
		fmt.Println(i)
		fmt.Printf("%c", v)
	} // 输出字符值
}
```

#### 字符串修稿

byte和rune类型

```go
package main

import "fmt"

func main() {
	var a = '我'
	var b = 'w'
	fmt.Println(string(a))
	fmt.Println(string(b))
}
```

Go语言的字符有两种：
- uint8类型，byte型，代表ASCII码的一个字符
- int32类型，rune类型，代表UTF-8类型

在UTF8编码下一个中文汉字由3～4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串

```go
package main

import "fmt"

func main() {
    s := "hello 世界!"
    b := []byte(s)
    b[5] = ','
    fmt.Printf("%s\n", s)
    fmt.Printf("%s\n", b)
    
    r := []rune(s)
    r[6] = '中'
    r[7] = '国'
    fmt.Println(s)
    fmt.Println(string(r))
}
```

### strings包

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
strs := "This is an example of a string, 世界！"
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "Th") 
	fmt.Printf("%t\n", strings.HasPrefix(strs, "Th")) // 前缀
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "string")
	fmt.Printf("%t\n", strings.HasSuffix(strs, "string")) //后缀
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "example")
	fmt.Printf("%t\n", strings.Contains(strs, "example")) // 包含
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "e y")
	fmt.Printf("%t\n", strings.ContainsAny(strs, "e y")) // 包含任意
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "世")
	fmt.Printf("%t\n", strings.ContainsRune(strs, '世')) // 包含字符
}
```
- strings.Contains substr为空，返回true
- strings.ContainsAny substr为空，返回false

#### 索引

如果找不到索引，返回-1

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hi, I'm Jusene.Hi,世界"
    
    fmt.Printf("The position of \"Jusene\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Jusene"))
    fmt.Printf("The position of the first instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Hi"))
    fmt.Printf("The position of the last instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
    fmt.Printf("The position of \"Tom\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Tom"))
    fmt.Printf("%d\n", strings.IndexRune(str, '世'))
    
}

```

#### 替换

strings.Replace(str, old, new, n) 一共四个参数，n表示匹配到第几个，n=-1表示匹配全部

```go
package main

import (
    "strings"
    "fmt"
)

func main() {
    str := "你好世界，这个世界真好。"
    new := "地球"
    old := "世界"
    n := 1
    fmt.Println(strings.Replace(str, old, new, n))
}
```

#### 统计

```go
package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func main() {
    str := "Goland is cool, right?"
    fmt.Printf("%d\n", strings.Count(str, "o"))
    fmt.Printf("%d\n", strings.Count(str, "oo"))
    
    stri := "你好世界"
    fmt.Printf("%d\n", len([]rune(stri)))
    fmt.Println(utf8.RuneCountInString(stri))
    
}
```

#### 大小写转换

```go
package main

import (
    "fmt"
    "strings"
) 

var origin string = "How are you! jusene"
var lower string
var upper string

func main() {
    fmt.Printf("%s\n", origin)
    lower = strings.ToLower(origin)
    fmt.Printf("%s\n", lower)
    upper = strings.ToUpper(origin)
    fmt.Printf("%s\n", upper)
}
```

#### 修剪

字符修剪

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Trim(" !!! Goland !!! ", " ! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Goland !!! ", "!"))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Goland !!! ", " ! "))
	fmt.Printf("%q\n", strings.TrimRight(" !!! Goland !!! ", " ! "))
	fmt.Println(strings.TrimSpace("\t\n 这是\t一句话\r\n\t"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今天"))
}
```

#### 分割

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	ls := strings.Split("A, B, C", ",")
	fmt.Printf("%s\n%s\n%s\n", ls[0], ls[1], ls[2])
}
```

#### 插入字符

strings.Join用于将元素类型string的slice使用分隔符拼接组成一个字符串。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the quick brower 中文"
	strSli := strings.Fields(str)
	fmt.Printf("%s\n", strSli)
	for _, val := range strSli {
		fmt.Printf("%s", val)
	}
	str2 := strings.Join(strSli, ";")
	fmt.Printf("%s\n", str2)

	str3 := strings.Split(str, " ")
	str4 := strings.Join(str3, ",")
	fmt.Println(str4)
}
```

strings.Filed 函数用于把字符串转换为字符串切片，通过range获得每个切片值

### strconv包

这个包用于字符串与其他类型的转换

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "123"
	fmt.Printf("orig当前是%T类型，且操作系统是%d位\n", orig, strconv.IntSize)

	// 字符转换int
	num, _ := strconv.Atoi(orig)
	fmt.Printf("%T\n", num)
	// 字符转换float64
	fl, _ := strconv.ParseFloat(orig, 64)
	fmt.Printf("%T\n", fl)
	// 十进制转换字符
	news := strconv.Itoa(num)
	fmt.Printf("%T\n", news)
	// 64位转换字符串
	// strconv.FormatFloat(f float64, fmt bytes, prce int, bitsize int) string
	new2 := strconv.FormatFloat(fl, 'f', 64, 64)
	fmt.Printf("%T\n", new2)
	// fmt表示格式('b','e','f','g'),prce表示精度，bitSize的值为32表示float32，64表示float64
}
```

### 字符格式化

- %% %字面量
- %b 一个二进制整数，将一个整数格式化为二进制的表达方式
- %c 一个Unicode的字符串
- %d 十进制数值
- %o 八进制数值
- %x 小写十六进制数值
- %X 大写十六进制数值
- %U 一个unicode表示法表示的整型码值，默认是4个整数字符
- %s 输出原生的UTF-8字节表示的字符，如果console不支持UTF-8编码，则会输出乱码
- %t 以true或者false的方式输出布尔值
- %v 使用默认格式输出值，或者如果方法存在，则使用类型String()方法输出自定义值
- %T 输出值的类型

### 强制类型转换

在强制类型转换时，需要注意数据长度被截断而发生的数据精度损失（如浮点型转换整型）与值溢出（值超过转换目标类型的值范围时）

```go
package main

import (
    "fmt"
)

func main() {
    a := 32
    b := 3
    fmt.Printf("%d", a/b)
    fmt.Printf("%f", float64(a)/float64(b))
}
```

#### 自定义类型

```go
package main

import (
    "fmt"
)

func main() {
    type stu struct {
        Name string
        Age int64
    }
    
    student := stu{
        Name: "guoxing",
        Age: 12,
    }
    
    fmt.Println(student.Name, student.Age)
}
```

#### 类型别名

byte类型是int8类型别名，rune类型实际上是int32类型别名

```go
package main

import (
    "fmt"
)

type (
    字符串 string
)

func main() {
    var b 字符串
    b = "这是中文"
    fmt.Println(b)
    a := "这也是中文"
    // fmt.Println(b+a) // 自定义类型别名和string是两个类型，不能直接相加
    fmt.Println(string(b)+a)
}
```

#### 指针

指针变量都是一个内存位置，每个内存位置都有其定义的地址，可以使用&运算

```go
package main 

import (
    "fmt"
)

func main() {
    a := 10
    fmt.Printf("%x\n", &a) 
}
```

指针是一个变量，其值是另一个变量的地址，所有指针的值的实际数据类型都是相同，他表示内存地址的长十六进制数

使用指针基本三个步骤：

- 定义一个指针变量
- 将一个变量的地址赋值给一个指针
- 最后访问指针变量中可用地址的值

```go
package main

import "fmt"

func main()  {
	a := 20
	ap := &a
	fmt.Printf("a的地址：%x\n", &a)
	fmt.Printf("ap的地址：%x\n", ap)
	fmt.Printf("*ap的地址：%x\n", *ap)
}
```

#### nil指针

```go
package main

import "fmt"

func main() {
	var pri *int
	fmt.Printf("%x\n", pri)
}
```

nil指针在标准库中定义值为0的常量

#### 指针的指针

```go
package main

import "fmt"

func main() {
	var null *int
	nu := &null
	fmt.Printf("null --> nil: %x\n", null)
	fmt.Printf("null的内存地址: %x\n", nu)
	fmt.Printf("nu --> null --> nil: %x\n", *nu)
	fmt.Printf("nu的内存地址：%x\n", &nu)


	A := 10
	AP := &A
	APP := &AP
	fmt.Printf("A: %d\n", A)
	fmt.Printf("AP: %x\n", AP)
	fmt.Printf("*AP: %d\n", *AP)
	fmt.Printf("APP: %x\n", APP)
	fmt.Printf("*APP: %x\n", *APP)
	fmt.Printf("**APP: %d\n", **APP)
}
```

#### 指针数组

```go
package main

import "fmt"

const MAX int = 3

func main() {

	a := []int{10, 100, 200}
	var ptr [MAX]*int

	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("a[%d]的地址 = %x\n", i, ptr[i] )
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d]的值是: %d\n", i, *ptr[i])
	}
}
```

#### 传递给函数

```go
package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf("交换之前a的值为: %d\n", a)
	fmt.Printf("交换之前b的值为: %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换之后a的值: %d\n", a)
	fmt.Printf("交换之后b的值: %d\n", b)
}

func swap(x *int, y *int) {
	var temp  int
	temp = *x
	*x = *y
	*y = temp

}
```
