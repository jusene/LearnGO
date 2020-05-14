package main

import "io/ioutil"

func main() {
	str := "hello world"
	data := []byte(str)
	ioutil.WriteFile("./src/github.com/jusene/day19/TEXT", data, 0644)
}
