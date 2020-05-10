package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	d, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(d))

	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "test agent")
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	d, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(d))
}
