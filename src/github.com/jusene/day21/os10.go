package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建一个硬链接
	err := os.Link("test", "oldtest")
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
