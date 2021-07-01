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
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
	// 修改切片内容以及map信息影响pc1
	pc2.SystemName = "MacOS"
	pc2.UseNumber = 100
	pc2.Memory.Count = 3
	pc2.Memory.MemorySize[0] = 8
	pc2.Fan["left"] = FanSpeed{200}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
}
