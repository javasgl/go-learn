/**
 * @author songgl
 * @since 2018-11-23 16:40
 */
package main

import (
	"log"
	"net"

	"grpc-example/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type server struct{}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello From " + in.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen :%v", err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Printf("listen on :%v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
