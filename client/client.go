package main

import (
	"fmt"
	"log"
	"rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := jsonrpc.Dial("udp", ":1234")
	if err != nil {
		log.Exit("dialing:", err)
	}
	log.Stdout("Connected")

	// Synchronous call
	args := &Args{7,8}
	var reply int
	for i := 0; i < 10; i++ {
		args.A = i
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Exit("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	}
}
