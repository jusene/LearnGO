package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	fmt.Println(string(data))
}
