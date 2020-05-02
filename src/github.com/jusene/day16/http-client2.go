package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{
		"name": []string{"jusene"},
		"age": []string{"27"},
	}
	data.Set("job", "enginer")
	u, err := url.ParseRequestURI("http://127.0.0.1:9090")
	if err != nil {
		fmt.Errorf("parse url requestUrl failed, err %v", err)
	}
	u.RawQuery = data.Encode()
	fmt.Printf("%s", u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(body))
}
