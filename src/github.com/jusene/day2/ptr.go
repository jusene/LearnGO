package main

import "fmt"

func main()  {
	a := 20
	ap := &a
	fmt.Printf("a的地址：%x\n", &a)
	fmt.Printf("ap的地址：%x\n", ap)
	fmt.Printf("*ap的地址：%x\n", *ap)
}
