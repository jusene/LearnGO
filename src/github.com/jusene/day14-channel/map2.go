package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i ++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.LoadOrStore(key, 100)
			fmt.Printf("k=%v, v=%v\n", key, value)
		}(i)
	}
	wg.Wait()
}
