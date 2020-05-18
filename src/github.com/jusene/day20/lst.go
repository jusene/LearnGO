package main

import (
	list2 "container/list"
	"fmt"
)

func main() {
	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf("长度: %v\n", list.Len())
	fmt.Printf("第一个元素: %#v, %#v\n", list.Front(), list.Front().Value)
	fmt.Printf("第二个元素: %#v, %#v\n", list.Front().Next(), list.Front().Next().Value)

}
