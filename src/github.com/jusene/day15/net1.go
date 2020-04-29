package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法：%s IP地址\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("无效IP地址")
	} else {
		fmt.Println("IP地址是", addr.String())
	}
}
