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

	//切片以及map新空间互不影响
	pc2.SystemName = "MacOs"
	pc2.UseNumber = 100
	pc2.Memory = Memory{Count: 2, MemorySize: []int{8, 8}}
	pc2.Fan = map[string]FanSpeed{"left": {2000}, "right": {1500}}
	pc2.Money = Money{1000.45}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
}
