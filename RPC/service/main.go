package main

import (
	. "RPC/model"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	err := rpc.Register(new(Arith))
	if err != nil {
		return
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Println("New Conn")
		go rpc.ServeConn(conn)
	}
	//go http.Serve(l, nil)
	//for {
	//	conn, err := l.Accept()
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println("New Conn")
	//	go func(conn net.Conn) {
	//		rpc.ServeConn(conn)
	//	}(conn)
	//}
}
