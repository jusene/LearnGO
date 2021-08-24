package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19-file/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	fmt.Println(string(buf[:n]))
}
