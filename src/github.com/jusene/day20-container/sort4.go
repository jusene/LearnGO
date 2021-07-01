package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 23, 22, 25, 20, 26, 122}
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println(s)
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//fmt.Println(s)
	fmt.Println(sort.SearchInts(s, 21))

	// 使用二分法算法来搜索某指定切片，如果slice以升序排序，则　f func中应该使用＞＝,如果slice以降序排序，则应该使用<=
	fmt.Println(sort.Search(len(s), func(i int) bool {
		return s[i] >= 10
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

func Search() {
	// 二分法
	var target int = 21
	num := []int{1, 23, 22, 25, 20, 26, 122}
	sort.Ints(num)
	fmt.Println(num)

	i, j := 0, len(num)
	f := func(h int) bool { return num[h] >= target }
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	fmt.Println(i)
}
