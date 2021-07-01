package main

import "fmt"

func main() {
	var user struct{Name string; Age int}
	user.Name = "JUSENE"
	user.Age = 27
	fmt.Printf("%#v\n", user)
}
