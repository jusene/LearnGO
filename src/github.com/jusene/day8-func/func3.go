package  main

import "fmt"

func main()  {
	ret := dosome()
	if ret == nil {
		fmt.Println(ret)
	}

}

func dosome() []int {
	return nil // nil可以看作一个有效的slice, 没必要显示返回一个长度为0的切片
}


