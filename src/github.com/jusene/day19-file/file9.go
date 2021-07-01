package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	if _, err := io.WriteString(fileobj, "hello"); err == nil {
		fmt.Println("ok")
	}
}
