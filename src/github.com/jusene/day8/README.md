## Go语言函数

```
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    // 逻辑代码
    return value1, value2 // 返回多值
}
```

函数名称如果小写开头，它的作用只属于所证明的包，不能被其他包调用；如果函数名以大写字母开头，则该函数便是公开的，可以被其他包调用。

Go语言函数不支持嵌套，重载和默认参数。

```go
package main

import "fmt"

func main() {
	x := 3
	y := 4
	a, b := SumAndProduct(x, y)
	fmt.Println(a, b)
	c, d := SumAndProduct1(x, y)
	fmt.Println(c, d)
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

// 命名的返回值就相当于在函数的声明的时候声明了一个变量
func SumAndProduct1(A, B int) (add int, mul int) {
	add = A + B
	mul = A * B
	return
}
```

### 函数作为参数

```go
package main

import "fmt"

func main() {
	num := 100
	filter(isBig, num)
}

type funcBool func(x int) bool

func filter(f funcBool, num int) bool {
	if f(num) {
		fmt.Println("ok")
		return true
	} else {
		fmt.Println("no")
		return false
	}
}

func isBig(x int) bool {
	if x >= 100 {
		return true
	}
	return false
}
```