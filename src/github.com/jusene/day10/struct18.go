package main

import "fmt"

type Aniaml struct {
	name string
}

func (a *Aniaml) move() {
	fmt.Printf("%s动了\n", a.name)
}

type Dog struct {
	Feet int8
	*Aniaml // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet:   4,
		Aniaml: &Aniaml{
			name: "葫芦",
		},
	}

	d1.move()
	d1.wang()
}
