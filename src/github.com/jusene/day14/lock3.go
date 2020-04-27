package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	defer wg.Done()
	rwlock.Lock() // 加写锁
	time.Sleep(time.Millisecond)
	rwlock.Unlock() // 解写锁
}

func read() {
	defer wg.Done()
	rwlock.RLock()  // 加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
}

func main() {
	start := time.Now()

	// 写锁并发时间18.9736ms
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}


	// 读锁并发时间1.0203ms
	/*
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	*/
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
