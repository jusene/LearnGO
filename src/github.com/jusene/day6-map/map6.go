package main

import "fmt"

func main() {
	m := make(map[string][]int)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Println(len(s), cap(s))
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	fmt.Println(s[:1], s[2:])
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
