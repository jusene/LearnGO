package main

import "fmt"

type Computer struct {
	CPU      string
	Memory   string
	HardDisk string
}

func (c *Computer) SetCPU(cpu string) {
	c.CPU = cpu
}

func (c *Computer) GetCPU() string {
	return c.CPU
}

func (c *Computer) SetMemory(memory string) {
	c.Memory = memory
}

func (c *Computer) GetMemory() string {
	return c.Memory
}

func (c *Computer) SetHardDisk(hardDisk string) {
	c.HardDisk = hardDisk
}

func (c *Computer) GetHardDisk() string {
	return c.HardDisk
}

// Builder规范建造者
type Builder interface {
	SetCPU(cpu string) Builder
	SetMemory(memory string) Builder
	SetHardDisk(hardDisk string) Builder
	Build() *Computer
}

// 根据不同的业务完成创建对象的组建
type ComputerBuilder struct {
	computer *Computer
}

func (c *ComputerBuilder) SetCPU(cpu string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetCPU(cpu)
	return c
}

func (c *ComputerBuilder) SetMemory(memory string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetMemory(memory)
	return c
}

func (c *ComputerBuilder) SetHardDisk(hardDisk string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetHardDisk(hardDisk)
	return c
}

func (c *ComputerBuilder) Build() *Computer {
	return c.computer
}

// 规范复杂对象的创建流程
type Director struct {
	Builder Builder
}

func (d Director) Create(cpu string, memory string, hardDisk string) *Computer {
	return d.Builder.SetCPU(cpu).SetMemory(memory).SetHardDisk(hardDisk).Build()
}

func main() {
	builder := new(ComputerBuilder)
	director := Director{Builder: builder}
	computer := director.Create("17", "32G", "4T")
	fmt.Println(*computer)
}
