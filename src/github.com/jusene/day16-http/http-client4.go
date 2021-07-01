package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 上传图片
func main() {
	data := url.Values{
		"name": []string{"jusene"},
		"age":  []string{"27"},
	}
	resp, err := http.PostForm("http://127.0.0.1:9090", data)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
