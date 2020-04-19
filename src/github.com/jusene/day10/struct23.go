package main

import "fmt"

type Camera struct {}

func (c *Camera) TakePicture() string {
	return "拍照"
}

type Phone struct {}

func (c * Camera) Call() string {
	return "响铃"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println(cp.TakePicture())
	fmt.Println(cp.Call())
}
