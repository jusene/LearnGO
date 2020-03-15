package main

import (
	"fmt"
	"strings"
)

func main() {
	ls := strings.Split("A, B, C", ",")
	fmt.Printf("%s\n%s\n%s\n", ls[0], ls[1], ls[2])
}
