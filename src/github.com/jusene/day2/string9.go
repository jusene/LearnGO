package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Trim(" !!! Goland !!! ", " ! "))
	fmt.Printf("%q\n", strings.Trim(" !!! Goland !!! ", "!"))
	fmt.Printf("%q\n", strings.TrimLeft(" !!! Goland !!! ", " ! "))
	fmt.Printf("%q\n", strings.TrimRight(" !!! Goland !!! ", " ! "))
	fmt.Println(strings.TrimSpace("\t\n 这是\t一句话\r\n\t"))
	fmt.Printf("%q\n", strings.Trim("今天天气真好", "今天"))
}