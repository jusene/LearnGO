package main

import (
	"log"
	"net"
	"time"
)

// TCP 服务端

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 5秒请求一次
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		log.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
	address := net.TCPAddr{
		IP: net.ParseIP("127.0.0.1"),
		Port: 8000,
	}
	listener, err := net.ListenTCP("tcp4", &address)
	checkError(err)
	*/

	address, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	checkError(err)
	listener, err := net.ListenTCP("tcp4", address)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		checkError(err)
		log.Println("远程地址：", conn.RemoteAddr())
		go echo(conn)
	}
}
