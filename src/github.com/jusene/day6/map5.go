package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do"
	fields := strings.Fields(str)

	var count = make(map[string]int)
	for _, v := range fields {
		value, ok := count[v]
		if ok {
			count[v] = value + 1
		} else {
			count[v] = 1
		}
	}
	fmt.Println(count)
}
