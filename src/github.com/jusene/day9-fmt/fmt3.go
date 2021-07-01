package main

import "fmt"

func main() {
	name := fmt.Sprint("jusene")
	age := 27
	n := fmt.Sprintf("name:%s age:%d", name, age)
	m := fmt.Sprintln("坚持")
	fmt.Println(n, m)
}
