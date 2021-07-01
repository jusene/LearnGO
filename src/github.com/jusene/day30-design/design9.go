package main

import "fmt"

// 策略模式
type Strategy interface {
	DoOperation(num1, num2 int) int
}

type OperationAdd struct {
}

func (o *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSub struct {
}

func (o *OperationSub) DoOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMul struct {
}

func (o *OperationMul) DoOperation(num1, num2 int) int {
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) Context(strategy Strategy) *Context {
	c.strategy = strategy
	return c
}

func (c *Context) Execute(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	context := new(Context).Context(new(OperationAdd))
	fmt.Println(context.Execute(10, 5))

	context = new(Context).Context(new(OperationSub))
	fmt.Println(context.Execute(10, 5))

	context = new(Context).Context(new(OperationMul))
	fmt.Println(context.Execute(10, 5))
}
