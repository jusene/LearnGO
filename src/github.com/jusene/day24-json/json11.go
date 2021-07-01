package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	c1 := Comment{Content: "永远保持谦逊"}
	i1 := Image{
		Title: "jusnee",
		URL:   "http://www.baidu.com",
	}

	// struct -> json string
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{
		&c1,
		&i1,
	})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	// json string -> struct
	jsonStr := `{"Content":"永远保持谦逊","title":"jusnee","url":"http://www.baidu.com"}`
	var (
		c2 Comment
		i2 Image
	)

	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}
