package main

import (
	"rpc/jsonrpc"
	"log"
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

type PacketListener struct {
	c net.PacketConn
	addr net.Addr
}

func (pl *PacketListener)  Read(b []byte) (n int, err os.Error) {
	n, pl.addr, err = pl.c.ReadFrom(b)
	return n, err
}

func (pl *PacketListener) Write(b []byte) (n int, err os.Error) {
	return pl.c.WriteTo(b, pl.addr)
}

func (pl *PacketListener) Close() os.Error {
	return pl.c.Close()
}

func serve(proto string, addr string) {
	switch proto {
	case "tcp":
		l, err := net.Listen(proto, addr)
		if err != nil {
			log.Exit(err)
		}
		for {
			conn, _ := l.Accept()
			jsonrpc.ServeConn(conn)
		}
	case "udp":
		pl := new(PacketListener)
		c, err := net.ListenPacket(proto, addr)
		if err != nil {
			log.Exit(err)
		}
		pl.c = c
		for {
			jsonrpc.ServeConn(pl)
		}
	default:
		log.Exit("Protocol", proto, "not supported")
	}

}

func main() {
	log.Stdout("Starting Server")
	arith := new(Arith)
	rpc.Register(arith)

	addr := ":1234"
	go serve("udp", addr)
	serve("tcp", addr)
}
