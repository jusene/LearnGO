package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "good study"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "good night"
	fmt.Println(msg)
	return msg
}

func printInterface(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	fmt.Println(v.Method(0).Type())
	fmt.Println(t.Method(0).Type)
	fmt.Println(t.Method(0).Name)
	fmt.Println(t.MethodByName("Sleep"))
	fmt.Println(v.MethodByName("Sleep").IsNil())
	for i := 0; i < v.NumMethod(); i ++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
		v.MethodByName("Sleep").Call(args)
		v.MethodByName("Study").Call(args)
	}
}


func main() {
	printInterface(student{
		Name:  "jusene",
		Score: 27,
	})
}
