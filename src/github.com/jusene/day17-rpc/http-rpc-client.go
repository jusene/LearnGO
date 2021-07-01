package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//serverAddress := os.Args[1]
	serverAddress := "127.0.0.1"

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Dial错误", err)
	}

	args := Args{17, 1}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith错误: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
