## Go语言包

### 包的结构

每个工作空间必须由bin, pkg, src三个目录组成， bin目录主要存放可执行文件，pkg目录存放编译好的库文件，主要*.a文件，src目录主要存放go的源代码。

定义包:
```go
package 包名
```

注意事项：

- 一个文件夹下只能有一个包，同样一个包的文件不能在多个文件下
- 包名可以不和文件夹的名字一样不能包含`-`符号
- 包名为main的包为应用程序的入口包，编译时不包含main的包的源代码时不会得到可执行文件

```go
package calc

import "fmt"

var name = "jusene"

func PrintCalc(x, y int) {
	fmt.Println(x + y)
}
```

```go
package main

import "github.com/jusene/day11/calc"

func main() {
	calc.PrintCalc(1, 2)
}
```

包的可见性，在Go语言中只需要将标识符的首字符大写就可以对外可见。

### 包的导入

- import导入语句通常放在文件开头包申明语句的下面
- 导入的包名需要使用双引号包裹起来
- 包名时从`$GOPATH/src/`后开始计算
- Go语言中禁止循环导入包

```
import "fmt"    // 对应/usr/local/go/pkg/linux_amd64/fmt.a
import "os/exec" // 对应/usr/local/go/pkg/linx_amd64/os/exec.a
```

### 导入声明

```
import "crypto/rand" // 默认模式 rand.Function
import R "crypto/rand" // 包重命名 R.Function
import . "crypto/rand" // 简便模式 Function
import _ "crypto/rand" // 匿名导入 仅让该包执行初始化函数
```

匿名导入执行初始化函数，init函数在包的是初始化函数，在导入的时候自动执行，init函数没有参数也没有返回值

包中init函数的执行时机：全局声明 -> init() -> main()

```go
package main

import (
    "database/sql"
    _ "github.com/lib/pq" // 启用对postgres的支持
    _ "github.com/go-sql-driver/mysql" // 启用对mysql的支持
)

db, err = sql.Open("postgres", dbname) // OK
db, err = sql.Open("mysql", dbname) // OK
db. err = sql.Open("sqlite3", dbname) // 返回错误：unknown driver "sqlite3"
```

### 导入路径

所有非标准库包的导入路径以所在组织的互联网域名为前缀，这样一来就有了一个独一无二的路径。

### Go语言工具箱

#### 包文档

根据编程规范，每个包都应该有包的注释，一般在package前的一个注释块。对于多文件包，包注释只需要出现在任意文件前即可。

#### 内部包

Go语言的封装只有两个特性，一个是私有的不可导出的成员，另一个是公开的可导出的成员。但是有些时候需要一种中间状态，对一些包公开，对一些包不可见，这就引入一个概念--内部包，在go语言中，构建工具会对导入路径包含internal关键字的包做特殊处理。一个internal包只能被和internal目录下同一个父目录的包导入。

例如：net/http/internal/chuncked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入。不过net/url包却可以导入net/http/httputil包

```
net/http
net/http/internal/chuncked
net/http/httputil
net/url
```

#### 查询包

go list 命令可以查询可用包的信息，其最简单的形式，可以测试包是否在工作区并打印它的导入路径。

```go
go list github.com/go-sql-driver/mysql
// 列出列表工作区中的所有包
go list ...
// 特定子目录下的所有包
go list github.com/...
// 某个主题相关的所有包
go list ...xml...
```

```go
go list -json .\github.com\jusene\day11\calc
{
        "Dir": "C:\\Users\\Jusene\\Desktop\\LearnGO\\src\\github.com\\jusene\\day11\\calc",
        "ImportPath": "github.com/jusene/day11/calc",
        "Name": "calc",
        "Target": "C:\\Users\\Jusene\\Desktop\\LearnGO\\pkg\\windows_amd64\\github.com\\jusene\\day11\\calc.a",
        "Root": "C:\\Users\\Jusene\\Desktop\\LearnGO",
        "Match": [
                "./github.com/jusene/day11/calc"
        ],
        "Stale": true,
        "StaleReason": "not installed but available in build cache",
        "GoFiles": [
                "pkg.go"
        ],
        "Imports": [
                "fmt"
        ],
        "Deps": [
                "errors",
                "fmt",
                "internal/bytealg",
                "internal/cpu",
                "internal/fmtsort",
                "internal/oserror",
                "internal/poll",
                "internal/race",
                "internal/reflectlite",
                "internal/syscall/windows",
                "internal/syscall/windows/registry",
                "internal/syscall/windows/sysdll",
                "internal/testlog",
                "io",
                "math",
                "math/bits",
                "os",
                "reflect",
                "runtime",
                "runtime/internal/atomic",
                "runtime/internal/math",
                "runtime/internal/sys",
                "sort",
                "strconv",
                "sync",
                "sync/atomic",
                "syscall",
                "time",
                "unicode",
                "unicode/utf16",
                "unicode/utf8",
                "unsafe"
        ]
}

```