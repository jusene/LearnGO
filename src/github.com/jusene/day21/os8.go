package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if os.IsPermission(err) {
		fmt.Println("文件没有写入权限")
	} else {
		log.Fatal(err)
	}

	defer file.Close()

	file, err = os.OpenFile("test", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if os.IsPermission(err) {
		fmt.Println("文件没有写入权限")
	} else {
		log.Fatal(err)
	}

	defer file.Close()
}
