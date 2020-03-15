package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf("交换之前a的值为: %d\n", a)
	fmt.Printf("交换之前b的值为: %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换之后a的值: %d\n", a)
	fmt.Printf("交换之后b的值: %d\n", b)
}

func swap(x *int, y *int) {
	var temp  int
	temp = *x
	*x = *y
	*y = temp

}
