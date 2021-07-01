package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Head("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Header["Server"][0])
}
