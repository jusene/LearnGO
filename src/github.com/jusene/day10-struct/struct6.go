package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct{}
	var b struct{User string; Age int}

	fmt.Println(unsafe.Sizeof(a)) // 空结构体不占用空间
	fmt.Println(unsafe.Sizeof(b))
}
