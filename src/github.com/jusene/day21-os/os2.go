package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	originaPath := "test.txt"
	newPath := "test2.txt"
	err = os.Rename(originaPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
