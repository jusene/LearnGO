package main

import "os"

func main() {
	file, err := os.OpenFile("test", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte("写入字节。\r\n"))

	file.WriteString("写入字符\r\n")

}
