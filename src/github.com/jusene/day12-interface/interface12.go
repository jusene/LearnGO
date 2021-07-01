package main

import "fmt"

func main() {
	var studInfo = make(map[string]interface{})
	studInfo["name"] = "jusene"
	studInfo["age"] = 27
	studInfo["marr"] = false

	fmt.Printf("%v", studInfo)
}
