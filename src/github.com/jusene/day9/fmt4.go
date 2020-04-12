package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
		}
	}()
	/*
		age, err := funErr()
		if err != nil {
			panic(err)
		}
		fmt.Print(age)
	*/
	age1, err1 := funcErr1()
	if err1 != nil {
		panic(err1)
	}
	fmt.Fprint(os.Stdout, age1)
}

func funErr() (age int, err error) {
	return 100, fmt.Errorf("错误")
}

func funcErr1() (age int, err error) {
	e := errors.New("new error")
	return 1, fmt.Errorf("%w", e)
}
