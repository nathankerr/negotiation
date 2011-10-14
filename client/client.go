package main

import (
	"fmt"
	"rpc/jsonrpc"
	"log"
//	"net"
	"os"
//	"crypto/rand"
	"rpc"
//	"time"
//	"crypto/tls"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//func dialTLS(addr string) (*rpc.Client, os.Error) {
//	conn, err := net.Dial("tcp", addr)
//	if err != nil {
//		return nil, err
//	}
//	config := &tls.Config{Rand: rand.Reader, Time: time.Nanoseconds}
//	ca := tls.NewCASet()
//	ca.SetFromPEM([]byte("ca.crt"))
//	tlsconn := tls.Client(conn, config)
//
//	return jsonrpc.NewClient(tlsconn), nil
//}

func dial(addr string, requirements string) (*rpc.Client, os.Error) {
	switch requirements {
	case "reliable":
		log.Println("tcp")
		return jsonrpc.Dial("tcp", addr)
//	case "secure":
//		log.Println("tls")
//		return dialTLS(addr)
	default:
		log.Println("udp")
		return jsonrpc.Dial("udp", addr)
	}
	return nil, os.EINVAL
}

func main() {
	client, err := dial(":1234", "")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	var reply int
	for i := 0; i < 10; i++ {
		args.A = i
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	}
}
