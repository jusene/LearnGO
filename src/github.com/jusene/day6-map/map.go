package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["张"] = 10
	dictMap["三"] = 20
	fmt.Println(dictMap)
	fmt.Println(dictMap["张"])

	dictMap1 := map[string]string{
		"username": "jusene",
		"password": "123rt",
	}
	fmt.Println(dictMap1["username"])

	dictMap2 := map[string][]int{}
	dictMap2["a"] = []int{1, 2, 3, 4}
	fmt.Println(dictMap2)
	fmt.Println(dictMap2["a"])

	// 不可以使用切片，函数以及包含切片结构类型由于具有引用语义，均不能作为映射的键
	// dictMap3 := map[[]string]int{}

	// value为map的切片
	dictMap4 := make([]map[string]int, 3)
	dictMap4[0] = make(map[string]int, 10)
	dictMap4[0]["name"] = 1
	dictMap4[0]["pass"] = 0
	fmt.Println(dictMap4)


}
