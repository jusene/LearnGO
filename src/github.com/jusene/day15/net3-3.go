package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// tcp 客户端

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法： %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.DialTimeout("tcp", service, 5 * time.Second)
	checkError(err)

	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		checkError(err)
		log.Println(strings.TrimSpace(string(data[0:n])))
	}
}
