package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

// Len() 人数
func (s StuScores) Len() int {
	return len(s)
}

// Less() 成绩将由低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap() 排序
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"张三", 95},
		{"李四", 91},
		{"赵六", 96},
		{"王六", 90},
	}

	fmt.Println("=======默认=======")
	// 原始顺序
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	fmt.Println()
	// stuScores已经实现了sort.Interface接口
	sort.Sort(stus)
	fmt.Println("=======排序之后=======")
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}

	fmt.Println(sort.IsSorted(stus))

	sort.Sort(sort.Reverse(stus)) // Reverse只是重新了Less方法
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}
}
