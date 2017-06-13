package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//server
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	//client
	args := &Args{7, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("airth error:", err)
	}
	fmt.Println("Arith:%d*%d=%d", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)

	replyCall := <-divCall.Done

	fmt.Println(replyCall)

}
