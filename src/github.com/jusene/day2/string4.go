package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := "This is an example of a string, 世界！"
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(strs, "Th")) // 前缀
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "string")
	fmt.Printf("%t\n", strings.HasSuffix(strs, "string")) //后缀
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "example")
	fmt.Printf("%t\n", strings.Contains(strs, "example")) // 包含
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "e y")
	fmt.Printf("%t\n", strings.ContainsAny(strs, "e y")) // 包含任意
	fmt.Printf("Does the string \"%s\" have prefix %s? \n", strs, "世")
	fmt.Printf("%t\n", strings.ContainsRune(strs, '世')) // 包含字符
}
