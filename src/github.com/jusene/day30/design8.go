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

//内存数量以及大小
type Memory struct {
	Count      int
	MemorySize []int
}

//电脑信息
type Computer struct {
	SystemName string              //系统名字
	UseNumber  int                 //使用次数
	Memory     Memory              //存储
	Fan        map[string]FanSpeed //风扇
	Money      Money               //售价
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
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

	// 浅拷贝
	pc2 := pc1.Clone()
	pc2.Memory.MemorySize[0] = 9999
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)

	// 深拷贝
	pc2 = pc1.BackUp()
	pc2.Memory.MemorySize[0] = 8888
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)
}
