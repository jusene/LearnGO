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
