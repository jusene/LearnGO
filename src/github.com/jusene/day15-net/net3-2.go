package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// tcp 客户端

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法： %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	TCPAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	checkError(err)

	conn.SetKeepAlive(true) // 会间隔性地发送keepalive包，操作系统可以通过该包来判断一个tcp连接是否已经断开

	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		checkError(err)
		log.Println(strings.TrimSpace(string(data[0:n])))
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%s", err.Error())
		os.Exit(1)
	}
}
