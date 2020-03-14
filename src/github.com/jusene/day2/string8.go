package main

import (
	"fmt"
	"strings"
)

var origin string = "How are you! jusene"
var lower string
var upper string

func main() {
	fmt.Printf("%s\n", origin)
	lower = strings.ToLower(origin)
	fmt.Printf("%s\n", lower)
	upper = strings.ToUpper(origin)
	fmt.Printf("%s\n", upper)
}
