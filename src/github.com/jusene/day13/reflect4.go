package main

import (
	"fmt"
	"reflect"
)

func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct {}{}
	// 尝试从结构体中查找abc字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 禅师从结构体中查找abc方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map不存在的键,", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")))
}
