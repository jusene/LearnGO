package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误", err.Error())
		os.Exit(2)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法：%s host:port", os.Args[0])
		os.Exit(2)
	}

	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	tick := time.Tick(5 * time.Second)
	for now := range tick {
		_ = now
		conn.Write([]byte("anything"))
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		checkError(err)
		fmt.Fprintf(os.Stdout, "%s\n", string(buf[0:n]))
	}
}
