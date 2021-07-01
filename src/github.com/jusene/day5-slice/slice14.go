package main

import "fmt"

func main() {
source := []int{10, 20, 30, 40, 50}
slice := source[2:3]
fmt.Println(len(slice), cap(slice))
// 容量足够，append会修改source的值
slice = append(slice, 11)
fmt.Println(slice, source)

slice1 := source[2:3:3]
// 容量不足， append会新创建一个新的底层数组
fmt.Println(len(slice1), cap(slice1))
slice2 := append(slice1, 11)
fmt.Println(slice2, source)
}
