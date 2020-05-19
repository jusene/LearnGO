package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("goDir", 0777)
	os.MkdirAll("goDir/test1/test2", 0777)
	err := os.Remove("goDir")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("goDir")
}
