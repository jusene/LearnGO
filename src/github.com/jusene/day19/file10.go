package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	writerobj := bufio.NewWriterSize(fileobj, 4096)

	if _, err = writerobj.WriteString("hello test"); err == nil {
		fmt.Println("ok")
	}

	buf := []byte("hello testets")
	if _, err := writerobj.Write(buf); err == nil {
		fmt.Println("ok")
		if err := writerobj.Flush(); err != nil {panic(err)}
		fmt.Println("flush")
	}

}
