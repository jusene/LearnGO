package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if readNum == 0 {
			break
		}

		fmt.Print(string(buf))
	}
}
