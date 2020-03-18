package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [...]int{3, 5, 4, -1, 9, 11, -14}
	slice := a[:]
	sort.Ints(slice)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)
}

