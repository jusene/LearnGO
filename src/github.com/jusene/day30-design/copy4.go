package main

import "fmt"

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

// 内存数量以及大小
type Memory struct {
	Count      int
	MemorySize []int
}

// 电脑信息
type Computer struct {
	SystemName string
	UseNumber  int
	Memory     Memory
	Fan        map[string]FanSpeed
	Money      Money
}

func main() {
	pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory: Memory{
			Count:      4,
			MemorySize: []int{32, 32, 32, 32},
		},
		Fan:   map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money: Money{123.45},
	}

	// 浅拷贝
	pc2 := pc1
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)

	Modify(pc2)
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)
}

func Modify(pc Computer) {
	fmt.Printf("pcinfo pc1:%v\n", pc)
	pc.SystemName = "MacOs"
	pc.UseNumber = 100
	pc.Memory.Count = 2
	pc.Memory.MemorySize[0] = 8
	pc.Memory.MemorySize[1] = 8
	pc.Memory.MemorySize[2] = 0
	pc.Memory.MemorySize[3] = 0
	pc.Fan["left"] = FanSpeed{2000}
	pc.Fan["right"] = FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v\n", pc)
}
