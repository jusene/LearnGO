package main

import (
	"fmt"
	"sync"
)

// 单例设计模式
// 双重检查机制
type SingleObject struct {
	Count int
}

var singleObj *SingleObject

var lock *sync.Mutex = &sync.Mutex{}

func GetInstance3() *SingleObject {
	if singleObj == nil {
		lock.Lock()
		defer lock.Unlock()
		singleObj = new(SingleObject)
	}
	return singleObj
}

func testSingleton1() {
	s1 := GetInstance3()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance3()
	fmt.Printf("%v, %v", &s2, s2)
}

func main() {
	testSingleton1()
}
