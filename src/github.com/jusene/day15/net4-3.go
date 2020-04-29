package main

import (
	"fmt"
	"net"
)

// UDP 服务端

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"),
		Port: 8000,
	})
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}

	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err: ", err)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%v", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
