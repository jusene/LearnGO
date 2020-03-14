package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a := 3.0
	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi)
	fmt.Println(reflect.TypeOf(a)) // 默认类型float64
}
