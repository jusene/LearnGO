package main

import "fmt"

// 单例设计模式
// 饿汉模式

type SingleObject struct {
	Count int
}

var singleObj *SingleObject

// 饿汉模式将在包加载的时候就创建单例对象，当程序中用不到该对象时，浪费了一部分空间
func init() {
	singleObj = new(SingleObject)
}

func GetInstance2() *SingleObject {
	return singleObj
}

func main() {
	s1 := GetInstance2()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance2()
	fmt.Printf("%v, %v", &s2, s2)
}
