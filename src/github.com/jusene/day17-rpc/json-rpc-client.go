package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	args := Args{15, 7}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
