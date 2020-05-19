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
