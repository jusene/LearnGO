package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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

// 基于序列化和反序列化来实现对象的深度拷贝
// 需要深拷贝的变量必须首字母大写才可以被拷贝
func deepcopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
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

	// 深拷贝
	pc2 := new(Computer)
	if err := deepcopy(pc2, pc1); err != nil {
		panic(err.Error())
	}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)

	pc2.SystemName = "MacOs"
	pc2.UseNumber = 100
	pc2.Memory.Count = 2
	pc2.Memory.MemorySize[0] = 8
	pc2.Memory.MemorySize[1] = 8
	pc2.Memory.MemorySize[2] = 0
	pc2.Memory.MemorySize[3] = 0
	pc2.Fan["left"] = FanSpeed{2000}
	pc2.Fan["right"] = FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
}
