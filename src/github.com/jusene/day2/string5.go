package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, I'm Jusene.Hi,世界"

	fmt.Printf("The position of \"Jusene\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Jusene"))
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"Tom\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Tom"))
	fmt.Printf("%d\n", strings.IndexRune(str, '世'))

}

