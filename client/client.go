package main

import (
	"fmt"
	"log"
	"os"
	"rpc"
	"rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func dial(addr string, requirements string) (*rpc.Client, os.Error) {
	switch requirements {
	case "reliable":
		log.Stdout("tcp")
		return jsonrpc.Dial("tcp", addr)
	default:
		log.Stdout("udp")
		return jsonrpc.Dial("udp", addr)
	}
	return nil, os.EINVAL
}

func main() {
	client, err := dial(":1234", "reliable")
	if err != nil {
		log.Exit("dialing:", err)
	}

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
