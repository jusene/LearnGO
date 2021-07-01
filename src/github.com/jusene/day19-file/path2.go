package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	rd, err := ioutil.ReadDir("C:\\")
	if err != nil {
		fmt.Println("read dir fail: ", err)
		return
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Println(fi.Name())
		} else {
			fmt.Println(filepath.Join("C:\\", fi.Name()))
		}
	}

	err = filepath.Walk("C:\\Users\\Jusene\\Downloads", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(info.IsDir(), info.Name())
		return nil
	})
}
