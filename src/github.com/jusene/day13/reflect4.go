package main

import (
	"fmt"
	"reflect"
)

func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// fmt.Println("nil IsNil:", reflect.ValueOf(nil).IsNil()) // 分类必须是通道、函数、接口、映射、指针、切片之一
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct {}{}
	// 尝试从结构体中查找abc字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 禅师从结构体中查找abc方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	// map
	c := map[string]int{
		"娜扎": 1,
	}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map存在的键,", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")))
	fmt.Println("map不存在的键,", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜胡")))
}
