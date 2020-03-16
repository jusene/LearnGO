package main

import "fmt"

func main() {
	i := [3]int{10, 20, 30}
	A(i)
	fmt.Println(i)

	j := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	B(j)
	fmt.Println(j)

	k := [2]int{10, 20}
	C(&k)
	fmt.Println(k)
}

func A(a [3]int) {
	a[0] = 100
}

func B(b [3][2]int) {
	b[2][0] = 100
}

func C(c *[2]int) {
	c[0] = 100
}