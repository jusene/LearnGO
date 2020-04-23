package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float64
	var b myInt
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age int
	}

	type book struct {
		title string
	}

	var d = person{
		name: "jusene",
		age:  27,
	}
	var e = book{title: "golang"}
	reflectType(d)
	reflectType(e)
}
