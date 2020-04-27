package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

func main() {

	for i, v := range make([]string, 10) {
		//once.Do(onces) // once只会被执行一次
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	wg.Wait()
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
