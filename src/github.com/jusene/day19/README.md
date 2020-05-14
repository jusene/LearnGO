## Go语言 I/O

Go语言将I/O操作封装在以下几个包:
- io 在io包中最重要的是两个接口-Reader和Writer接口，只要程序实现了这两个接口，就有I/O的功能
- io/ioutil 封装一些实用的I/0函数，这个包提供一些常用的I/O操作函数
- fmt 实现格式化I/O
- bufio 实现带缓存的I/O，封装于io.Reader和io.Writer对象

### io包

io 实现两个接口，它就有了I/O的功能

1. Reader接口

```
type Reader interface {
   Read(p []byte) (n int, err error) 
}
```

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	data, err := ReadFrom(os.Stdin, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	data, err := ReadFrom(strings.NewReader("from string"), 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```

2. Writer 接口

```
func Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Stringer 接口

```go
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age int
	Sex int
}

func (p *Person) String() string {
	buffer := bytes.NewBufferString("这是")
	buffer.WriteString(p.Name + ",")
	if p.Sex == 0 {
		buffer.WriteString("他")
	} else {
		buffer.WriteString("她")
	}
	buffer.WriteString("今年")
	buffer.WriteString(strconv.Itoa(p.Age))
	buffer.WriteString("岁。 ")
	return buffer.String()
}

func main() {
	person := &Person{
		Name: "jusene",
		Age:  27,
		Sex:  0,
	}

	fmt.Println(Stringer(person).String())
}
```

### 文件处理

- vbytes: byte slice便利操作，在Go语言中string是内置类型，同样它与普通的slice类型有相似的性质。
- strconv: 字符串和基本数据类型之间的装换
- regexp: 正则表达式
- unicode: Unicode编码

```
file, err := os.Open("file.go")
if err != nil {
    panic(err)
}
```

读取文件
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	fmt.Println(string(buf[:n]))
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("%s", string(buf))
	}
}
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	fmt.Println(string(data))
}
```

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		fmt.Println(lineText)
	}
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if readNum == 0 {
			break
		}

		fmt.Print(string(buf))
	}
}
```

写文件

```go
package main

import "io/ioutil"

func main() {
	str := "hello world"
	data := []byte(str)
	ioutil.WriteFile("./src/github.com/jusene/day19/TEXT", data, 0644)
}
```

```go
package main

import "os"

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	_, err = fileobj.WriteString("hello golang")
	if err != nil {
		panic(err)
	}
}
```

```
const (
        O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
        O_WRONLY int = syscall.O_WRONLY // 只写打开文件
        O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
        O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
        O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
        O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
        O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
        O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
)
```

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	if _, err := io.WriteString(fileobj, "hello"); err == nil {
		fmt.Println("ok")
	}
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	writerobj := bufio.NewWriterSize(fileobj, 4096)

	if _, err = writerobj.WriteString("hello test"); err == nil {
		fmt.Println("ok")
	}

	buf := []byte("hello testets")
	if _, err := writerobj.Write(buf); err == nil {
		fmt.Println("ok")
		if err := writerobj.Flush(); err != nil {panic(err)}
		fmt.Println("flush")
	}
}
```