package main

import (
	"http"
	"rpc/jsonrpc"
	"log"
	"net"
	"os"
	"crypto/rand"
	"rpc"
	"crypto/tls"
	"time"
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
	c    net.PacketConn
	addr net.Addr
}

func (pl *PacketListener) Read(b []byte) (n int, err os.Error) {
	n, pl.addr, err = pl.c.ReadFrom(b)
	return n, err
}

func (pl *PacketListener) Write(b []byte) (n int, err os.Error) {
	return pl.c.WriteTo(b, pl.addr)
}

func (pl *PacketListener) Close() os.Error {
	return pl.c.Close()
}

// rpc over tcp
func serveTCP(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Exit(err)
	}
	defer l.Close()
	for {
		conn, _ := l.Accept()
		jsonrpc.ServeConn(conn)
	}
}

// rpc over udp
func serveUDP(addr string) {
	pl := new(PacketListener)
	c, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Exit(err)
	}
	defer pl.Close()
	pl.c = c

	for {
		jsonrpc.ServeConn(pl)
	}
}

// rpc over tls over tcp
func serveTLS(addr string) {
	config := &tls.Config{
		Rand: rand.Reader,
		Time: time.Nanoseconds,
	}
	config.Certificates = make([]tls.Certificate, 1)
	var err os.Error
	config.Certificates[0], err = tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Exit(err)
	}

	l, err := tls.Listen("tcp", ":1235", config)
	if err != nil {
		log.Exit(err)
	}
	for {
		conn, _ := l.Accept()
		jsonrpc.ServeConn(conn)
	}
}

// rpc over HTTP
func serveHTTP(addr string) {
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Exit(err)
	}
	http.Serve(l, nil)
}

func main() {
	log.Stdout("Starting Server")
	arith := new(Arith)
	rpc.Register(arith)

	go serveUDP(":1234")
	go serveTCP(":1234")
	go serveTLS(":1235")
	serveHTTP(":1236")
}
