package main

import "fmt"

// 面试题 下面的代码可以通过编译吗？

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = Student{} // 接受者为指针，*Student{}可以通过编译
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
