package main

import (
	"errors"
	"fmt"
)

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "*":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x * y
}

func main() {
	op, err	:= do("+")
	if err != nil {
		panic(err)
	}
	ret := op(10, 20)
	fmt.Println(ret)
}
