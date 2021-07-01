package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["jusene"] = 90
	dictMap["zgx"] = 200

	for k, v := range dictMap {
		fmt.Println(k, v)
	}
}
