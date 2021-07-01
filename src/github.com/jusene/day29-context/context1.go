package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	defer wg.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现goroutine
	wg.Wait()
	fmt.Println("over")
}
