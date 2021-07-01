package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		fmt.Println(lineText)
	}

}
