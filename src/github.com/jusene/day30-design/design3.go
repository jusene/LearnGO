package main

import (
	"fmt"
)

// 单例设计模式
// 懒汉模式
type SingleObject struct {
	Count int
}

var singleObj *SingleObject

// 存在线程安全问题，多线程时，会创建多个对象，所有出现饿汉模式
func GetInstance1() *SingleObject {
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

func testSingleton() {
	s1 := GetInstance1()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance1()
	fmt.Printf("%v, %v", &s2, s2)
}

func main() {
	testSingleton()
}
