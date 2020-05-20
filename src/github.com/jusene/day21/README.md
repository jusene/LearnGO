## Go语言 文件处理

- 创建目录/删除目录

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("goDir", 0777)
	os.MkdirAll("goDir/test1/test2", 0777)
	err := os.Remove("goDir")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("goDir")
}
```

- 创建文件和查看状态

```go
package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	fileInfo os.FileInfo
	err error
	path = "test1/test2/"
	fileName = "demo"
	filePath = path + fileName
)

func main() {
	// 创建文件夹
	err = os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("成功创建目录")
	}

	// 创建空白的文件
	newFile, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newFile)
	newFile.WriteString("HELLO WORLD")
	newFile.Close()

	// 查看文件的信息，如果文件不存在，则返回错误
	fileInfo, err = os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		log.Fatal("文件不存在")
	}

	fmt.Println("文件名称: ", fileInfo.Name())
	fmt.Println("文件大小: ", fileInfo.Size())
	fmt.Println("文件权限: ", fileInfo.Mode())
	fmt.Println("最后修改时间: ", fileInfo.ModTime())
	fmt.Println("是否为文件夹: ", fileInfo.IsDir())
	fmt.Printf("系统接口类型: %T", fileInfo.Sys())
	fmt.Printf("系统信息: %+v\n", fileInfo.Sys())
}
```

- 重命名与移动

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	originaPath := "test.txt"
	newPath := "test2.txt"
	err = os.Rename(originaPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
```

- 打开与关闭

```go
package main

import (
	"log"
	"os"
)

var (
	file *os.File
	fileInfo os.FileInfo
	err error
	dirPath = "test1/test2/"
	fileName = "demo"
	filePath = dirPath + fileName
)

func main() {
	os.MkdirAll(dirPath, 0777)
	file, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// 打印文件内容
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()

	file, err = os.OpenFile(dirPath, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
```

- 删除与截断

```go
package main

import "os"

func main() {
	file, _ := os.Create("test")
	file.WriteString("hello world")
	file.Close()
	// err := os.Remove("test")
	err := os.Truncate("test", 1) // 传入0会清空文件
	if err != nil {
		panic(err)
	}
}
```

- 复制文件

```go
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 创建文件
	file, err := os.Create("test")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("HELLO GOLANG")
	file.Close()

	// 打开文件
	originFile, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}

	defer originFile.Close()

	// 创建新的文件复制文件内容到新文件
	newFile, err := os.Create("test_copy")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// 从源文件中复制到新文件
	bytesWritten, err := io.Copy(newFile, originFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bytesWritten)

	// 将文件内容Flush到硬盘
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
```

- 跳转函数

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test")
	if err != nil {
		panic(err)
	}
	file.WriteString("hello golang")
	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5

	// 用来计算offset的初始位置
	// 0 文件开始位置
	// 1 当前位置
	// 2 文件结尾数
	whence := 0

	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		panic(err)
	}

	fmt.Println("移到位置5: ", newPosition)

	// 从当前位置回退两字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("回退两个字节: ", newPosition)

	// 获取当前位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("当前: ", currentPosition)

	// 移到文件开头
	newPosition, err = file.Seek(0, 0)
	fmt.Println("开头", newPosition)
}
```

- 写入函数

```go
package main

import "os"

func main() {
	file, err := os.OpenFile("test", os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte("写入字节。\r\n"))

	file.WriteString("写入字符\r\n")	
}
```

- 权限控制

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil && os.IsNotExist(err){
		fmt.Println("文件不存在")
	} else if os.IsPermission(err) {
		fmt.Println("文件没有写入权限")
	} else {
		log.Fatal(err)
	}

	defer file.Close()

	file, err = os.OpenFile("test", os.O_RDONLY | os.O_CREATE, 0666)
	if err != nil && os.IsNotExist(err){
		fmt.Println("文件不存在")
	} else if os.IsPermission(err) {
		fmt.Println("文件没有写入权限")
	} else {
		log.Fatal(err)
	}

	defer file.Close()
}
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	err := os.Chmod("test", 0777)
	if err != nil {
		log.Println(err)
	}

	// 改变文件所有者
	err = os.Chown("test", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// 查看文件信息
	fileInfo, err := os.Stat("test")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("文件不存在")
		}
		log.Fatal(err)
	}

	fmt.Println(fileInfo.ModTime())

	// 改变时间戳
	twoDayFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDayFromNow
	lastModifyTime := twoDayFromNow
	err = os.Chtimes("test", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(fileInfo.ModTime())
}
```

- 文件链接

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建一个硬链接
	err := os.Link("test", "oldtest" )
	if err != nil {
		panic(err)
	}
	
	// 创建一个软链接
	err = os.Symlink("test", "softtest")
	if err != nil {
		panic(err)
	}
	
	// Lstat返回文件信息，如果文件是软链接，返回软链接的信息
	softinfo, err := os.Lstat("softtest")
	fmt.Printf("%v", softinfo)
	
	// 改变软链接的拥有者不会影响原始文件
	err = os.Lchown("softtest", os.Getuid(), os.Getgid())
}
```