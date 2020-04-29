package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// tcp 客户端

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	checkError(err)
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputinfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputinfo) == "Q" {
			return
		}

		_, err = conn.Write([]byte(inputinfo)) // 发送数据
		checkError(err)

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		checkError(err)

		fmt.Println(string(buf[:n]))
	}
}
