package main

import "os"

func main() {
	fileobj, err := os.OpenFile("./src/github.com/jusene/day19/TEXT", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()
	_, err = fileobj.WriteString("hello golang")
	if err != nil {
		panic(err)
	}
}
