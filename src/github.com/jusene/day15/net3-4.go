package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp 服务端

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[0:])
		checkError(err)
		recvStr := string(buf[:n])
		fmt.Println("接受client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	checkError(err)

	for {
		conn, err := listen.Accept()
		checkError(err)
		go process(conn)
	}
}
