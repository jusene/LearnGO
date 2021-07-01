package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	data, err := ReadFrom(strings.NewReader("from string"), 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
