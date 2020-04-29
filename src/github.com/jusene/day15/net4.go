package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// UDP 服务端

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误", err.Error())
		os.Exit(2)
	}
}

func main() {
	service := "127.0.0.1:1200"
	udpaddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp4", udpaddr)
	checkError(err)

	for {
		buf := make([]byte, 512)
		n, addr, err := conn.ReadFromUDP(buf)
		checkError(err)
		fmt.Fprintf(os.Stdout, "接受%s\n", addr.IP)
		fmt.Fprintf(os.Stdout, "%s\n", string(buf[0:n]))
		daytime := time.Now().String()
		conn.WriteToUDP([]byte(daytime), addr)
	}
}
