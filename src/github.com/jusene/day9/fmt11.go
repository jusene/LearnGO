package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bufioDemo()
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
