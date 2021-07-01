package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "test1/test2/"
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
