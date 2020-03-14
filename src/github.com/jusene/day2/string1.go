package main

import "fmt"

var str string = "武汉加油"

func main()  {
	for i := 0; i < len(str); i++ {
		fmt.Println(i)
		fmt.Printf("%c", str[i])
	} // 输出字节码

	for i, v := range str {
		fmt.Println(i)
		fmt.Printf("%c", v)
	} // 输出字符值
}