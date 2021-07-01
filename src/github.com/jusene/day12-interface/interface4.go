package main

import "fmt"

type Mover interface {
	move()
}

type dog struct{}

func (d dog) move() {
	fmt.Println("狗动了")
}

type cat struct{}

func (c *cat) move() {
	fmt.Println("猫动了")
}

func main() {
	wangcai := dog{} //值接受者实现接口
	fugui := &dog{}
	Mover(wangcai).move()
	Mover(fugui).move() // Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

	//miao := cat{} // 值接受者实现接口
	aiai := &cat{}
	//Mover(miao).move() // Mover不接受miao的类型
	Mover(aiai).move() // Mover可以接受*cat类型

}
