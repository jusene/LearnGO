package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("./src/github.com/jusene/day19/reader.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
