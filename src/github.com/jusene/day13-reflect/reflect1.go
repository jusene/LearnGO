package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	t := reflect.ValueOf(x)
	fmt.Printf("type:%v value:%v\n", v, t)
}

func main() {
	var a float64
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}
