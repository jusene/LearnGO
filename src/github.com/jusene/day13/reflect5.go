package main

import (
	"fmt"
	"reflect"
)

func refelectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%T,%v\n", v, v)
	k := v.Kind() // 获取底层的数据类型
	switch k {
	case reflect.Float32:
		fmt.Printf("%T\n", v.Float())
		ret := float32(v.Float())
		fmt.Printf("%T, %v", ret, ret)
	case reflect.Int32:
		fmt.Printf("%T\n", v.Int())
		ret := int32(v.Int())
		fmt.Printf("%T, %v", ret, ret)
	}
}

func main() {
	var aa int32 = 100
	refelectValue(aa) // 通过反射获得接口的值
}
