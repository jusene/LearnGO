package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	dictMap  := make(map[string]int, 200)
	for i := 0; i< 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)

		dictMap[key] = value
	}

	keys := make([]string, 0, 200)
	for key := range dictMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, dictMap[key])
	}
}
