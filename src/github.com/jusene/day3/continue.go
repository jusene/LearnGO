package main

import "fmt"

func main() {
LOOP1:
	for i := 0;i <= 5;i++ {
		switch {
		case i > 0:
			fmt.Println("A")
			continue LOOP1
		case i == 1:
			fmt.Println("B")
		default:
			fmt.Println("C")
		}
		fmt.Printf("i is: %d\n", i)
	}
}
