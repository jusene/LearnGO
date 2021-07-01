package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["jusene"] = 90
	dictMap["zgx"] = 100

	delete(dictMap, "jusene")
	fmt.Println(dictMap)
}
