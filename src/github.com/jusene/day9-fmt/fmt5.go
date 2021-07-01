package main

import (
	"fmt"
)

func main() {
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%")
}
