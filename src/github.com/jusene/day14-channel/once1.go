package main

import (
	"fmt"
	"sync"
	"time"
)

type singleton struct {
	name string
	age int
}

var instance *singleton
var once sync.Once
var wg sync.WaitGroup
var lock sync.Mutex

func GetInstance() *singleton {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("load instance...")
		instance = &singleton{
			name: "jusene",
			age: 27,
		}
		fmt.Println(instance.name, instance.age)
	})
	lock.Lock()
	fmt.Println("*** 过了一年")
	time.Sleep(time.Second)
	instance.age += 1
	lock.Unlock()
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetInstance()
	}
	wg.Wait()

	fmt.Println(instance.name, instance.age)
}
