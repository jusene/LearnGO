package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct {
	name string
}

func (d dryer) dry() {
	fmt.Printf("%s甩甩\n", d.name)
}

type haier struct {
	dryer // 嵌入甩干机
}

func (h haier) wash() {
	fmt.Printf("%s洗洗\n", h.dryer.name)
}

func main() {
	var machine WashingMachine
	h := haier{dryer{name: "海尔"}}
	machine = h
	machine.dry()
	machine.wash()
}
