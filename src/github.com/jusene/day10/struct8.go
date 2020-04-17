package main

import (
	"fmt"
	"reflect"
)

type tagType struct {
	goods bool "是否有存货"
	name string "商品名称"
	price float64 "商品价格"
}

func main() {
	tt := tagType{
		goods: true,
		name:  "IPHONE SE",
		price: 1000,
	}

	for i := 0; i < 3 ; i++ {
		refTag(tt, i)
	}
}

func refTag(tt tagType, ix int) {
	ttType := reflect.TypeOf(tt)
	iField := ttType.Field(ix)
	fmt.Printf("%v\n", iField.Tag)
}
