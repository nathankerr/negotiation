package main

import (
	"log"
	"http"
	"net"
	"os"
	"rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) os.Error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) os.Error {
	if args.B == 0 {
		return os.ErrorString("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func serve(proto string, addr string) {
	var l net.Listener
	var e os.Error

	switch proto {
	case "tcp":
		l, e = net.Listen(proto, addr)
	/*case "udp":
		l, e := net.ListenUDP(proto, addr)
		listener = net.Listener(l)*/
	default:
		log.Exit("Protocol ", proto, " not supported")
	}
	if e != nil {
		log.Exit("listen error:", e)
	}
	http.Serve(l, nil)
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	addr := ":1234"
	go serve("tcp", addr)
	serve("udp", addr)
}
