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

type UDPListener struct {
	c net.Conn
}

func (l *UDPListener) Accept() (c net.Conn, err os.Error) {
	if l == nil || l.c == nil {
		return nil, os.EINVAL
	}
	return l.c, nil
}

func (l *UDPListener) Close() os.Error {
	return l.c.Close()
}

func (l *UDPListener) Addr() net.Addr {
	return l.c.LocalAddr()
}

func ListenUDP(proto string, laddr string) (l *UDPListener, err os.Error) {
	l = new(UDPListener)
	var la *net.UDPAddr
	if laddr != "" {
		if la, err = net.ResolveUDPAddr(laddr); err != nil {
			return nil, err
		}
	}
	c, err := net.ListenUDP(proto, la)
	if err != nil {
		return nil, err
	}
	l.c = c
	return l, nil
}

func serve(proto string, addr string) {
	var l net.Listener
	var e os.Error

	switch proto {
	case "tcp":
		l, e = net.Listen(proto, addr)
	case "udp":
		l, e = ListenUDP(proto, addr)
	default:
		log.Exit("Protocol", proto, "not supported")
	}
	if e != nil {
		log.Exit("listen error:", e)
	}
	http.Serve(l, nil)
}

func main() {
	log.Stdout("Starting Server")
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	addr := ":1234"
	serve("udp", addr)
}
