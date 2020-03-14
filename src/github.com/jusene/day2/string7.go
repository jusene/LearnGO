package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Goland is cool, right?"
	fmt.Printf("%d\n", strings.Count(str, "o"))
	fmt.Printf("%d\n", strings.Count(str, "oo"))

	stri := "你好世界"
	fmt.Printf("%d\n", len([]rune(stri)))
	fmt.Println(utf8.RuneCountInString(stri))

}
