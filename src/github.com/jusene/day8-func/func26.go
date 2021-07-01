package main

import "fmt"

// n个台阶， 一次可以走1步，也可以走2步，有多少种走法

func main()  {
	fmt.Println(taijie(3))
}

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n - 1) + taijie(n - 2)
}