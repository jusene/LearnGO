
package main

import "fmt"

func main() {
	s := "hello 世界!"
	b := []byte(s)
	b[5] = ','
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)

	r := []rune(s)
	r[6] = '中'
	r[7] = '国'
	fmt.Println(s)
	fmt.Println(string(r))
}
