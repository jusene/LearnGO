package main

import "fmt"


func main() {
	var v1 complex64
	v1 = 3.2 + 12i
	v2 := 2.3 + 12i
	v3 := complex(3.2, 12)
	v := v2 + v3
	fmt.Println(v1, v2, v3, v)
}
