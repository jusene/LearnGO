package main

import "fmt"

func main() {
	var d bool
	justifyType(d)
}

func justifyType(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unsupport type")
	}
}