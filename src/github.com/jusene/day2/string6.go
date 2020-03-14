package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "你好世界，这个世界真好。"
	new := "地球"
	old := "世界"
	n := 1
	fmt.Println(strings.Replace(str, old, new, n))
}
