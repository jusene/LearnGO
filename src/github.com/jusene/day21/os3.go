package main

import (
	"log"
	"os"
)

var (
	file     *os.File
	fileInfo os.FileInfo
	err      error
	dirPath  = "test1/test2/"
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
