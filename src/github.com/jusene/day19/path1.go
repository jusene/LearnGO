package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "/home/example/example.go"
	dir := filepath.Dir(path)
	fmt.Println(dir)
	base := filepath.Base(path)
	fmt.Println(base)

	path = "/home/exmaple/"
	dir = filepath.Dir(path)
	fmt.Println(dir)
	base = filepath.Base(path)
	fmt.Println(base)

	path = "/home/exmaple"
	dir = filepath.Dir(path)
	fmt.Println(dir)
	base = filepath.Base(path)
	fmt.Println(base)

	path = "/home/example/example.go"
	winpath := "C:\\Pram\\exam\\example.go"
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs(winpath))
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Getwd())
	// 获取相对路径
	fmt.Println(filepath.Rel("/home/example", "/home/example/src/example.go"))
	fmt.Println(filepath.Rel("/home/exmaple", "/home/src/exmaple.go"))

	// 路径分隔
	fmt.Println(filepath.Split("/home/example/exmaple.go"))
	fmt.Println(filepath.Split("/home/example"))
	fmt.Println(filepath.Split("/home/example/"))
	fmt.Println(filepath.Split("example.go"))

	// 路径拼接
	fmt.Println(filepath.Join("/home", "example"))
	fmt.Println(filepath.SplitList("/example/exmple.go;/hello/world.go"))

	// 最短路径
	fmt.Println(filepath.Clean("/../../exmaple/example.go"))

	// 文件路径匹配
	fmt.Println(filepath.Match("/*", "/CI"))
	fmt.Println(filepath.Match("?", ""))
	fmt.Println(filepath.Match("[^123]", "5"))
	fmt.Println(filepath.Match("[12]", "1"))

	fmt.Println(filepath.Glob("C:\\*"))

}
