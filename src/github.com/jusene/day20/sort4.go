package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 21, 3, 4, 0, 10}
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	fmt.Println(sort.SearchInts(s, 21))

	// 使用二分法算法来搜索某指定切片，如果slice以升序排序，则　f func中应该使用＞＝,如果slice以降序排序，则应该使用<=
	fmt.Println(sort.Search(len(s), func(i int) bool {
		return s[i] <= 1
	}))

	a := []string{"a", "b", "g", "d"}
	sort.Strings(a)
	fmt.Println(sort.StringsAreSorted(a))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	fmt.Println(a)
	fmt.Println(sort.SearchStrings(a, "g"))

	f := []float64{1.22, 2.22, 0.12}
	sort.Float64s(f)
	fmt.Println(sort.Float64sAreSorted(f))
	fmt.Println(f)
	sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	fmt.Println(f)
	fmt.Println(sort.SearchFloat64s(f, 2.22))

}
