package main

import (
	"fmt"
	"sort"
)

type num []int

func (s num) Len() int {
	return len(s)
}

func (s num) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s num) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	x := 11
	n := num{12, 22, 31, 413, 11}

	sort.Sort(n)
	pos := sort.Search(len(n), func(i int) bool {
		return n[i] >= x
	})

	if pos < len(n) && n[pos] == x {
		fmt.Println(x, "在n的位置为：", pos)
	} else {
		fmt.Println("n不包含元素", x)
	}
}
