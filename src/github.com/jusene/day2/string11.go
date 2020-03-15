package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the quick brower 中文"
	strSli := strings.Fields(str)
	fmt.Printf("%s\n", strSli)
	for _, val := range strSli {
		fmt.Printf("%s", val)
	}
	str2 := strings.Join(strSli, ";")
	fmt.Printf("%s\n", str2)
	str3 := strings.Split(str, " ")
	str4 := strings.Join(str3, ",")
	fmt.Println(str4)
}