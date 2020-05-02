package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:9090/post"

	// form
	 // contentType := "application/x-www-form-urlencoded"
	 // data := "name=jusene&age=27"
	// json
	contentType := "application/json"
	data := `{"name": "jusene", "age": 27}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed; err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}
