package main

import (
	"fmt"
	"strings"
)

func checkFile(filename string) func(prefix, suffix string) bool {
	return func(prefix, suffix string) bool {
		if strings.HasPrefix(filename, prefix) && strings.HasSuffix(filename, suffix) {
			return true
		}
		return false
	}
}

func main() {
	jpgFunc := checkFile("test.jpg")
	fmt.Println(jpgFunc("test", "jpg"))
	fmt.Println(jpgFunc("test", "txt"))
}
